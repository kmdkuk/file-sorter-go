package cmd

import (
	"bufio"
	"github.com/kmdkuk/file-sorter-go/cmd/option"
	"github.com/kmdkuk/file-sorter-go/log"
	"github.com/spf13/cobra"
	"io"
	"os"
	"sort"
)

func init(){
	rootCmd.Flags().StringVarP(&option.Opt.InputFile, "input", "i", option.Opt.InputFile,"sort target file")
	rootCmd.Flags().StringVarP(&option.Opt.OutputFile, "output", "o", option.Opt.OutputFile, "sorted file")
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		log.Error("error: ", err)
		log.Info(rootCmd.UsageString())
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use: "file-sorter",
	Short: "This sorts the lines in the file.",
	Long: "This sorts the lines in the file.",
	SilenceErrors: true,
	SilenceUsage: true,
	Run: run,
}

func run(cmd *cobra.Command, _ []string){
	log.Debug("option.Opt: ",option.Opt)
	log.Debug("InputFile: ", option.Opt.InputFile)
	log.Debug("OutputFile: ", option.Opt.OutputFile)
	inputFile, err := open(option.Opt.InputFile)
	if err != nil{
		log.Error("error: ", err)
		log.Info(cmd.UsageString())
		os.Exit(1)
	}
	defer inputFile.Close()

	text, _ := read(inputFile)
	log.Debug("before: ", text)
	sort.Strings(text)
	log.Debug("after: ", text)
	if option.Opt.InputFile == option.Opt.OutputFile && option.Opt.OutputFile != option.STDOUT {
		log.Info("originalファイルを.originalとして保存します.")
		if err := makeCopyOriginal(inputFile); err != nil{
			log.Error("error: ", err)
			os.Exit(1)
		}
	}

	outputFile, err := open(option.Opt.OutputFile)
	if err != nil {
		log.Error("error: ", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	err = writeText(outputFile, text)
	if err != nil {
		log.Error("error: ", err)
		os.Exit(1)
	}
}

func open(filename string) (*os.File, error){
	if filename == option.STDIN {
		return os.Stdin, nil
	}
	if filename == option.STDOUT {
		return os.Stdout, nil
	}
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func read(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	output := make([]string, 1)
	for line:=0; scanner.Scan();line++ {
		if len(output) < line + 1 {
			diff := line - len(output)
			output = append(output, make([]string, diff + 1)...)
		}
		output[line] = scanner.Text()
	}
	return output, nil
}

func writeText(output io.Writer, text []string) error{
	for line:=0; line < len(text);line++ {
		_, err := output.Write([]byte(text[line]+"\n"))
		if err != nil{
			return err
		}
	}
	return nil
}

func makeCopyOriginal(file *os.File) error {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil{
		return err
	}
	dst, err := os.Create(file.Name()+".original")
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}
	return nil
}