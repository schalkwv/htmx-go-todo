package main

import "net/http"

func handleGetTasks(w http.ResponseWriter, _ *http.Request) {
	items, err := fetchTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := fetchCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	completedCount, err := fetchCompletedCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks := Tasks{
		Items:          items,
		Count:          count,
		CompletedCount: completedCount,
	}
	tmpl.ExecuteTemplate(w, "Base", tasks)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		tmpl.ExecuteTemplate(w, "Form", nil)
		return
	}
	_, err := insertTask(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err := fetchCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	tmpl.ExecuteTemplate(w, "Form", nil)
	tmpl.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})

}
