package word2txt

import (
	"bufio"
	"os"

	"github.com/unidoc/unioffice/document"
)

//Converter for extract data from word document
type Converter struct {
	Src       string
	Dst       string
	Content   []string
	IsWindows bool
}

//Extract data  from word doc
func (c *Converter) Extract(filename string) error {
	c.Src = filename
	doc, err := document.Open(c.Src)
	if err != nil {
		return err
	}
	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
		var line string
		for _, r := range p.Runs() {
			line += r.Text()
		}
		c.Content = append(c.Content, line)
	}
	return nil
}

//Output to dst
func (c *Converter) Output() error {
	f, err := os.OpenFile(c.Dst, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewWriter(f)

	if c.IsWindows {
		for _, p := range c.Content {
			br.WriteString(p)
			br.WriteString("\n")
		}
	} else {
		for _, p := range c.Content {
			br.WriteString(p)
			br.WriteString("\r\n")
		}
	}

	return f.Close()
}
