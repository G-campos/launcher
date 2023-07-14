package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Launcher App")
	myWindow.Resize(fyne.NewSize(400, 200))

	// Criar uma imagem para a animação
	image := canvas.NewImageFromFile("./carbuddy.jpg")

	// Configurar o widget da imagem
	image.FillMode = canvas.ImageFillOriginal

	// Criar um widget de container para posicionar a imagem no centro
	container := fyne.NewContainerWithLayout(layout.NewCenterLayout(), image)

	// Adicionar o widget de container à janela
	myWindow.SetContent(container)

	// Mostrar a janela
	myWindow.Show()

	// Aguardar um pouco para exibir a animação
	time.Sleep(time.Second * 3)

	// Remover o widget da imagem
	myWindow.SetContent(widget.NewLabel("Aplicativo iniciado!"))

	time.AfterFunc(3*time.Second, myWindow.Close)

	myApp.Run()

}
