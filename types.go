package main

type Document struct {
	Date string
	Ticker string
	High int
	Low int
	Open int
	Close int
}

func createDocument(Date string, Ticker string, High int, Low int, Open int, Close int) *Document {
	return &Document{
		Date: Date,
		Ticker: Ticker,
		High: High,
		Low: Low,
		Open: Open,
		Close: Close,
	}
}