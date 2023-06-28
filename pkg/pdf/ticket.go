package pdf

import (
	"errors"
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func CreatePDF(pdfFileName, heading string, logo string, data [][]string) error {
	maroto := pdf.NewMaroto(consts.Portrait, consts.A4)
	maroto.SetPageMargins(20, 10, 20)
	if err := buildHeading(maroto, heading, logo); err != nil {
		return err
	}
	if err := buildTicketDataList(maroto, data); err != nil {
		return err
	}
	if err := maroto.OutputFileAndClose(pdfFileName); err != nil {
		return errors.New("failed to save the ticket as PDF file")
	}
	return nil
}

func buildHeading(m pdf.Maroto, heading string, imageLogo string) (err error) {
	m.RegisterHeader(func() {
		m.Row(40, func() {
			m.Col(12, func() {
				err = m.FileImage(imageLogo, props.Rect{
					Center:  true,
					Percent: 75,
				})
			})
		})
	})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to load '%s' image in PDF", imageLogo))
	}
	m.Row(20, func() {
		m.Col(12, func() {
			m.QrCode(heading, props.Rect{
				Left:    0,
				Top:     5,
				Center:  true,
				Percent: 200,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(heading, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	return nil
}

func buildTicketDataList(m pdf.Maroto, data [][]string) error {
	purpleColor := getPurpleColor()
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Airplane Ticket", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList([]string{"", ""}, data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{6, 6},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{6, 6},
		},
		Align:                consts.Left,
		AlternatedBackground: &purpleColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})
	return nil
}

func getPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getGreyColor() color.Color {
	return color.Color{
		Red:   206,
		Green: 206,
		Blue:  206,
	}
}
