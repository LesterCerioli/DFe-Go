package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type MainWindow struct {
	model         *CTeTesteModel
	window        *gtk.Window
	webXmlEnvio   *gtk.Widget
	webXmlRetorno *gtk.Widget
}

func NewMainWindow() *MainWindow {
	mw := &MainWindow{}
	mw.model = NewCTeTesteModel()
	mw.model.SucessoSync = mw.Sucesso

	// Initialize GTK window and widgets here
	// ...

	return mw
}

func (mw *MainWindow) Sucesso(e RetornoEEnvio) {
	glib.IdleAdd(func() {
		mw.XmlTemp(e.Envio, "envio-tmp.xml")
		mw.XmlTemp(e.Retorno, "retorno-tmp.xml")
		mw.webXmlEnvio.LoadURI(filepath.Join(mw.getPath(), "envio-tmp.xml"))
		mw.webXmlRetorno.LoadURI(filepath.Join(mw.getPath(), "retorno-tmp.xml"))
	})
}

func (mw *MainWindow) XmlTemp(xml, nomeXml string) {
	err := ioutil.WriteFile(filepath.Join(mw.getPath(), nomeXml), []byte(xml), 0644)
	if err != nil {
		fmt.Println("Error writing XML file:", err)
	}
}

func (mw *MainWindow) getPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	return filepath.Dir(filename)
}

func (mw *MainWindow) Certificado_Click() {
	mw.model.ObterSerialCertificado()
}

func (mw *MainWindow) ArquivoCertificado_Click() {
	mw.model.ObterCertificadoArquivo()
}

func (mw *MainWindow) BuscarDiretorioSchema_Click() {
	mw.model.BuscarDiretorioSchema()
}

func (mw *MainWindow) BuscarDiretorioSalvarXml_Click() {
	mw.model.BuscarDiretorioSalvarXml()
}

func (mw *MainWindow) SalvarConfiguracoesXml_Click() {
	mw.model.SalvarConfiguracoesXml()
}

func (mw *MainWindow) ConsultarStatusServico() {
	mw.model.ConsultarStatusServico2()
}

func (mw *MainWindow) InutilizacaoDeNumeracao_Click() {
	mw.model.InutilizacaoDeNumeracao()
}

func (mw *MainWindow) ConsultaPorProtocolo_Click() {
	mw.model.ConsultaPorProtocolo()
}

func (mw *MainWindow) EventoCancelarCTe_Click() {
	mw.model.EventoCancelarCTe()
}

func (mw *MainWindow) EventoDesacordoCTe_Click() {
	mw.model.EventoDesacordoCTe()
}

func (mw *MainWindow) CartaCorrecao_Click() {
	mw.model.CartaCorrecao()
}

func (mw *MainWindow) CriarEnviarCTe2_Click() {
	mw.model.CriarEnviarCTe2e3()
}

func (mw *MainWindow) MainWindow_OnLoaded() {
	mw.model.CarregarConfiguracoes()
}

func (mw *MainWindow) ConsultaPorNumeroRecibo_Click() {
	mw.model.ConsultaPorNumeroRecibo()
}

func (mw *MainWindow) CriarEnviarAutomaticoCTe2_Click() {
	mw.model.CriarEnviarCTeConsultaReciboAutomatico2e3()
}

func (mw *MainWindow) CTeDistribuicaoDFe_Click() {
	mw.model.DistribuicaoDFe()
}

func (mw *MainWindow) EmitirCteOs_Click() {
	mw.model.EmitirCteOs()
}

func (mw *MainWindow) LoadXmlCte_Click() {
	dialog, err := gtk.FileChooserDialogNewWith2Buttons(
		"Open File",
		mw.window,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Cancel", gtk.RESPONSE_CANCEL,
		"Open", gtk.RESPONSE_ACCEPT)

	if err != nil {
		fmt.Println("Error creating file chooser dialog:", err)
		return
	}
	defer dialog.Destroy()

	filter, err := gtk.FileFilterNew()
	if err != nil {
		fmt.Println("Error creating file filter:", err)
		return
	}
	filter.AddPattern("*.xml")
	filter.SetName("XML files")
	dialog.AddFilter(filter)

	response := dialog.Run()
	if response == gtk.RESPONSE_ACCEPT {
		filename := dialog.GetFilename()
		if filename == "" {
			gtk.MessageDialogNew(mw.window, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Selecione um xml").Run()
		} else {
			mw.model.LoadXmlCTe(filename)
		}
	}
}

func main() {
	gtk.Init(nil)

	win := NewMainWindow()
	win.window.ShowAll()

	gtk.Main()
}
