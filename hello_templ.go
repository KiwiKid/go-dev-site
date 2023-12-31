// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func hello(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"flex items-center justify-center h-screen text-center font-medium font-mono\"><div class=\"font-bold rounded-lg border shadow-lg p-10 space-y-3\"><h1 class=\"text-4xl font-medium p-5\">")
		if err != nil {
			return err
		}
		var_2 := `Hey, i&apos;m Greg`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><h2 class=\"text-2xl font-medium\">")
		if err != nil {
			return err
		}
		var_3 := `&#123; a software developer, football player, hobby fpv drone pilot`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_4 := `&#125;`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><h2>")
		if err != nil {
			return err
		}
		var_5 := `&#123; based in christchurch, new zealand &#125; `
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><div class=\"text-lg border border-black p-4 rounded-md\"><div class=\"py-2\">")
		if err != nil {
			return err
		}
		var_6 := `open to new software development opportunities`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div>")
		if err != nil {
			return err
		}
		var_7 := `I thrive in fast paced, high trust environments, with a clear path`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_8 := `to having a big impact and everyone moving in the same direction.`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div class=\"py-2\">")
		if err != nil {
			return err
		}
		var_9 := `only considering fully/mostly remote roles`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div><div class=\"grid grid-cols-1 sm:grid-cols-2\"><a href=\"/dev/nzcovidmap\"><img src=\"/icons/covid19/icon-512x512.png\" alt=\"Locations of Interest Explorer Icon\" width=\"100\" height=\"100\"> ")
		if err != nil {
			return err
		}
		var_10 := `NZCovidMap`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><a href=\"/about\">")
		if err != nil {
			return err
		}
		var_11 := `About`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></div></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
