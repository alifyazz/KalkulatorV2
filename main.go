package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CalculatorData struct {
	Result     string
	Expression string
	History    []HistoryItem
}

type HistoryItem struct {
	Expression string
	Result     string
	Timestamp  time.Time
}

var history []HistoryItem

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/calculate", handleCalculate)
	http.HandleFunc("/clear", handleClear)
	http.HandleFunc("/history", handleHistory)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := CalculatorData{
		Result:     "0",
		Expression: "",
		History:    history,
	}
	tmpl.Execute(w, data)
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	expression := r.FormValue("expression")
	result := calculateExpression(expression)

	// Add to history if calculation was successful
	if result != "Error" && result != "0" {
		historyItem := HistoryItem{
			Expression: expression,
			Result:     result,
			Timestamp:  time.Now(),
		}
		history = append(history, historyItem)

		// Keep only last 10 items in history
		if len(history) > 10 {
			history = history[len(history)-10:]
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := CalculatorData{
		Result:     result,
		Expression: expression,
		History:    history,
	}
	tmpl.Execute(w, data)
}

func handleClear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Just clear the display, don't save to history
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := CalculatorData{
		Result:     "0",
		Expression: "",
		History:    history,
	}
	tmpl.Execute(w, data)
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func calculateExpression(expr string) string {
	if expr == "" {
		return "0"
	}

	// Clean expression
	expr = strings.ReplaceAll(expr, "×", "*")
	expr = strings.ReplaceAll(expr, "÷", "/")

	result, err := evaluateExpression(expr)
	if err != nil {
		return "Error"
	}

	return strconv.FormatFloat(result, 'f', -1, 64)
}

func evaluateExpression(expr string) (float64, error) {
	// Handle scientific functions first
	if strings.HasPrefix(expr, "sin") {
		return handleSin(expr[3:])
	}
	if strings.HasPrefix(expr, "cos") {
		return handleCos(expr[3:])
	}
	if strings.HasPrefix(expr, "tan") {
		return handleTan(expr[3:])
	}
	if strings.HasPrefix(expr, "log") {
		return handleLog(expr[3:])
	}
	if strings.HasPrefix(expr, "ln") {
		return handleLn(expr[2:])
	}
	if strings.HasPrefix(expr, "sqrt") {
		return handleSqrt(expr[4:])
	}
	if strings.HasPrefix(expr, "pow") {
		return handlePow(expr[3:])
	}

	// Handle basic arithmetic operations
	return evaluateBasicExpression(expr)
}

func handleSin(expr string) (float64, error) {
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}
	return math.Sin(num * math.Pi / 180), nil // Convert degrees to radians
}

func handleCos(expr string) (float64, error) {
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}
	return math.Cos(num * math.Pi / 180), nil
}

func handleTan(expr string) (float64, error) {
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}
	return math.Tan(num * math.Pi / 180), nil
}

func handleLog(expr string) (float64, error) {
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}
	return math.Log10(num), nil
}

func handleLn(expr string) (float64, error) {
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}
	return math.Log(num), nil
}

func handleSqrt(expr string) (float64, error) {
	num, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(num), nil
}

func handlePow(expr string) (float64, error) {
	// Simple power function - expects format like "2,3" for 2^3
	parts := strings.Split(expr, ",")
	if len(parts) != 2 {
		return 0, nil
	}
	base, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}
	exponent, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, err
	}
	return math.Pow(base, exponent), nil
}

func evaluateBasicExpression(expr string) (float64, error) {
	// Simple expression evaluator for basic operations
	// This handles expressions like "2+3*4" in a basic way

	// First, try to parse as a single number
	if num, err := strconv.ParseFloat(expr, 64); err == nil {
		return num, nil
	}

	// Handle basic operations with simple regex
	// This is a simplified version - for production use a proper parser

	// Addition
	if strings.Contains(expr, "+") {
		parts := strings.Split(expr, "+")
		if len(parts) == 2 {
			a, err1 := strconv.ParseFloat(parts[0], 64)
			b, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 == nil && err2 == nil {
				return a + b, nil
			}
		}
	}

	// Subtraction
	if strings.Contains(expr, "-") {
		parts := strings.Split(expr, "-")
		if len(parts) == 2 {
			a, err1 := strconv.ParseFloat(parts[0], 64)
			b, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 == nil && err2 == nil {
				return a - b, nil
			}
		}
	}

	// Multiplication
	if strings.Contains(expr, "*") {
		parts := strings.Split(expr, "*")
		if len(parts) == 2 {
			a, err1 := strconv.ParseFloat(parts[0], 64)
			b, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 == nil && err2 == nil {
				return a * b, nil
			}
		}
	}

	// Division
	if strings.Contains(expr, "/") {
		parts := strings.Split(expr, "/")
		if len(parts) == 2 {
			a, err1 := strconv.ParseFloat(parts[0], 64)
			b, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 == nil && err2 == nil {
				if b == 0 {
					return 0, nil // Return 0 for division by zero
				}
				return a / b, nil
			}
		}
	}

	// If we can't parse it, return 0
	return 0, nil
}
