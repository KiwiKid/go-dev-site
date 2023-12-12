// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func nzCovidMapWithContainer() templ.Component {
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
		_, err = templBuffer.WriteString("<html>")
		if err != nil {
			return err
		}
		err = headerComponent("Greg C. 👨‍💻 ✈️ ⚽").Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<body class=\"bg-gray-100\">")
		if err != nil {
			return err
		}
		err = nzCovidMap().Render(ctx, templBuffer)
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

func nzCovidMap() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"container mx-auto p-4\"><!--")
		if err != nil {
			return err
		}
		var_3 := ` Title and Introduction `
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"text-center\"><h1 class=\"text-4xl font-bold mb-2\">")
		if err != nil {
			return err
		}
		var_4 := `NZCovidMap`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"text-lg mb-4\">")
		if err != nil {
			return err
		}
		var_5 := `Over the course of the pandemic, the NZCovidMap reached over 200,000 Kiwis across New Zealand, with a peak of 15,000 daily unique visitors.`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div><!--")
		if err != nil {
			return err
		}
		var_6 := ` Project Description `
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"mb-6\"><p class=\"mb-2\">")
		if err != nil {
			return err
		}
		var_7 := `This tool provided people with a mobile-friendly view of Covid-19 Locations of Interest published by the Ministry of Health. As locations are found, the application automatically maintained Reddit posts for any new locations in relevant community subreddits.`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p class=\"mb-2\">")
		if err != nil {
			return err
		}
		var_8 := `A creative circle-based location selection method allows the user to view the information for many locations in a specific area.`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div><!--")
		if err != nil {
			return err
		}
		var_9 := ` Goals Section `
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"mb-6\"><h2 class=\"text-3xl font-semibold mb-2\">")
		if err != nil {
			return err
		}
		var_10 := `Goals`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><ul class=\"list-disc pl-5\"><li>")
		if err != nil {
			return err
		}
		var_11 := `Provide a Mobile-friendly view of Locations of Interest.`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li><li>")
		if err != nil {
			return err
		}
		var_12 := `Raise awareness of Locations of Interests in our community.`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li><li>")
		if err != nil {
			return err
		}
		var_13 := `Lower anxiety for those seeking updates about their community.`
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li></ul></div><!--")
		if err != nil {
			return err
		}
		var_14 := ` Tech Section `
		_, err = templBuffer.WriteString(var_14)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"mb-6\"><h2 class=\"text-3xl font-semibold mb-2\">")
		if err != nil {
			return err
		}
		var_15 := `Tech`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><ul class=\"list-disc pl-5\"><li>")
		if err != nil {
			return err
		}
		var_16 := `A Mobile-friendly map view for digesting localized information.`
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li><li>")
		if err != nil {
			return err
		}
		var_17 := `Up to date social Media sharing icons created for each town/city.`
		_, err = templBuffer.WriteString(var_17)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li><li>")
		if err != nil {
			return err
		}
		var_18 := `Reddit Integration - Auto-Generate a "Location of Interest summary" for each subreddit & share relevant subreddit.`
		_, err = templBuffer.WriteString(var_18)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li></ul></div><!--")
		if err != nil {
			return err
		}
		var_19 := ` 1.0 Design Section `
		_, err = templBuffer.WriteString(var_19)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"mb-6\"><h2 class=\"text-3xl font-semibold mb-2\">")
		if err != nil {
			return err
		}
		var_20 := `1.0 Design`
		_, err = templBuffer.WriteString(var_20)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><ol class=\"list-decimal pl-5\"><li>")
		if err != nil {
			return err
		}
		var_21 := `An n8n workflow that reads the CSV file published by the MoH. Locations are then categorized and saved.`
		_, err = templBuffer.WriteString(var_21)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li><li>")
		if err != nil {
			return err
		}
		var_22 := `An n8n workflow endpoint that returns all the locations in the last 30 days.`
		_, err = templBuffer.WriteString(var_22)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li><li>")
		if err != nil {
			return err
		}
		var_23 := `A statically rendered Next.JS site that is published each hour. A mobile-friendly interactive map with all locations marked.`
		_, err = templBuffer.WriteString(var_23)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li></ol></div><!--")
		if err != nil {
			return err
		}
		var_24 := ` 2.0 Design Section `
		_, err = templBuffer.WriteString(var_24)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"mb-6\"><h2 class=\"text-3xl font-semibold mb-2\">")
		if err != nil {
			return err
		}
		var_25 := `2.0 Design ▲`
		_, err = templBuffer.WriteString(var_25)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p>")
		if err != nil {
			return err
		}
		var_26 := `To reduce the complexity, cost and on-going maintenance cost of the application, the backend was deprecated in favor of directly pulling information from the MoH API. Recent improvements include generation of location-specific meta screenshots for engaging social media preview images.`
		_, err = templBuffer.WriteString(var_26)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div><!--")
		if err != nil {
			return err
		}
		var_27 := ` Image Placeholders `
		_, err = templBuffer.WriteString(var_27)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("--><div class=\"grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6\"><div class=\"rounded overflow-hidden shadow-lg\"><img class=\"w-full\" src=\"/path/to/your/image1.jpg\" alt=\"Image Description 1\"></div><div class=\"rounded overflow-hidden shadow-lg\"><img class=\"w-full\" src=\"/path/to/your/image2.jpg\" alt=\"Image Description 2\"></div><div class=\"rounded overflow-hidden shadow-lg\"><img class=\"w-full\" src=\"/path/to/your/image3.jpg\" alt=\"Image Description 3\"></div></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
