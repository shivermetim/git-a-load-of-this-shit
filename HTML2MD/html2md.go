package md "github.com/JohannesKaufmann/html-to-markdown"
import "github.com/JohannesKaufmann/html-to-markdown/plugin"

converter.Use(plugin.GitHubFlavored())
converter := md.NewConverter("", true, nil)

html = `<strong>Important</strong>`

markdown, err := converter.ConvertString(html)
if err != nil {
  log.Fatal(err)
}
fmt.Println("md ->", markdown)