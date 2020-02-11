package main

import (
        "fmt"
        "golang.org/x/net/context"
        "google.golang.org/api/option"
        "google.golang.org/api/sheets/v4"
        "log"
    )


const (
    client_secret_path = "./credentials/client_secret.json"
)

func main() {
    s, _ := NewSpreadsheetService()
    pushRequest := SpreadsheetPushRequest{
        SpreadsheetId: "INSERT_YOUR_ID",
        Range: "INSERT_YOUR_RANGE",
        Values: []interface{}{"One", "Two", "Three", "Four", "Five"},
    }
    s.WriteToSpreadsheet(&pushRequest)
}

type SpreadsheetService struct {
    service *sheets.Service
}

func NewSpreadsheetService() (*SpreadsheetService, error) {
    // Service account based oauth2 two legged integration
    ctx := context.Background()
    srv, err := sheets.NewService(ctx, option.WithCredentialsFile(client_secret_path), option.WithScopes(sheets.SpreadsheetsScope))

    if err != nil {
        log.Fatalf("Unable to retrieve Sheets Client %v", err)
    }

    c := &SpreadsheetService{
        service: srv,
    }

    return c, nil
}


func (s *SpreadsheetService) WriteToSpreadsheet(object *SpreadsheetPushRequest) error {
    var vr sheets.ValueRange
    vr.Values = append(vr.Values, object.Values)

    res, err := s.service.Spreadsheets.Values.Append(object.SpreadsheetId, object.Range, &vr).ValueInputOption("RAW").Do()

    fmt.Println("spreadsheet push ", res)

    if err != nil {
        fmt.Println("Unable to update data to sheet ", err)
    }

    return err
}   

type SpreadsheetPushRequest struct {
    SpreadsheetId string        `json:"spreadsheet_id"`
    Range         string        `json:"range"`
    Values        []interface{} `json:"values"`
}