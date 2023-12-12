// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "time"
import "strconv"

var currentYear = time.Now().Year()
var yearsOfExperience = currentYear - 2014
var yearsOfExperienceStr = strconv.Itoa(yearsOfExperience)

func aboutPage() templ.Component {
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
		_, err = templBuffer.WriteString("<div><div class=\"w-full text-lg\"><div class=\"m-auto sm:w-4/5 max-w-5xl\"><div class=\"w-full text-center\">")
		if err != nil {
			return err
		}
		var_2 := `🚨This page is a work in progress...🚨`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		var_3 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"w-full text-center\"><p class=\"italic\">")
			if err != nil {
				return err
			}
			var_4 := `(As a Kiwi, this level of self-promotion makes me deeply uncomfortable, but here we go...)`
			_, err = templBuffer.WriteString(var_4)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p></div> <p>")
			if err != nil {
				return err
			}
			var_5 := `I have been developing in a professional capacity as a full stack developer for over `
			_, err = templBuffer.WriteString(var_5)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var var_6 string = yearsOfExperienceStr
			_, err = templBuffer.WriteString(templ.EscapeString(var_6))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_7 := `years.`
			_, err = templBuffer.WriteString(var_7)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p>")
			if err != nil {
				return err
			}
			var_8 := `I have extensive experience in implementing complex software solutions, based on varying degrees of specification.`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p>")
			if err != nil {
				return err
			}
			var_9 := `Working closely with stakeholders, refining and architecting solutions through to deployment, maintenance and support.`
			_, err = templBuffer.WriteString(var_9)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p>")
			if err != nil {
				return err
			}
			var_10 := `I have a passionate about solving real problems, working with the best tools (or getting there) and improving the overall effectiveness of the organisation and the team`
			_, err = templBuffer.WriteString(var_10)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p>")
			if err != nil {
				return err
			}
			var_11 := `I am excited about continuing to exploring the benefit and opportunities of serveless/cloud solutions.`
			_, err = templBuffer.WriteString(var_11)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <div class=\"text-center w-full\"><a class=\"underline text-2xl\" href=\"")
			if err != nil {
				return err
			}
			var var_12 templ.SafeURL = `mailto:${settings.contactEmail}`
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_12)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\">")
			if err != nil {
				return err
			}
			var_13 := `Get in touch`
			_, err = templBuffer.WriteString(var_13)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = aboutSection("Introduction").Render(templ.WithChildren(ctx, var_3), templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<img src=\"/img/atopthehill.jpg\" alt=\"Myself on the top of a hill after climbing 1.4 vertical km to retrieve a fallen drone...\" width=\"5184\" height=\"3880\">")
		if err != nil {
			return err
		}
		var_14 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"grid grid-cols-3 space-y-2 sm:grid-cols-3 sm:space-x-4\"><div><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_15 := `Loving`
			_, err = templBuffer.WriteString(var_15)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_16 := `` + "`" + `go` + "`" + `, htmx, Node-Red, Tailwind`
			_, err = templBuffer.WriteString(var_16)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div><div><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_17 := `Learning`
			_, err = templBuffer.WriteString(var_17)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_18 := `MQTT, Django, postgresML, Unraid, htmx, tailwind`
			_, err = templBuffer.WriteString(var_18)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div><div><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_19 := `Up Next`
			_, err = templBuffer.WriteString(var_19)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_20 := `Alpine.js, Nix/NixOS`
			_, err = templBuffer.WriteString(var_20)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div><div class=\"col-span-full\"><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_21 := `Experience In`
			_, err = templBuffer.WriteString(var_21)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_22 := `C#, .Net Core + DI, React, Next.js + Vercel, SQL, Entity Framework, SQL Server, T-SQL, ASP.NET MVC, Web APIs, IIS, AWS CDK, RESTful API design, git, CI/CD, AWS Lambda, n8n, Storybook,DynamoDB, Heroku, AutoFac, Moq, MQTT, Swagger`
			_, err = templBuffer.WriteString(var_22)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = aboutSection("Tech").Render(templ.WithChildren(ctx, var_14), templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func aboutSection(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_23 := templ.GetChildren(ctx)
		if var_23 == nil {
			var_23 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div><h2>")
		if err != nil {
			return err
		}
		var var_24 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_24))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2>")
		if err != nil {
			return err
		}
		err = var_23.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func aboutWithContainer() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_25 := templ.GetChildren(ctx)
		if var_25 == nil {
			var_25 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html>")
		if err != nil {
			return err
		}
		err = headerComponent("About me").Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<body>")
		if err != nil {
			return err
		}
		err = aboutPage().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

type settingsType struct {
	contactEmail string
}
