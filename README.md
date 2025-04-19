<div id="top">

<!-- HEADER STYLE: CLASSIC -->
<div align="center">

# GO-CRUD

<em>Efficient CRUD Operations at Your Fingertips!</em>

<!-- BADGES -->
<img src="https://img.shields.io/github/license/PlumpAlbert/go-crud?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
<img src="https://img.shields.io/github/last-commit/PlumpAlbert/go-crud?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/PlumpAlbert/go-crud?style=default&color=0080ff" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/PlumpAlbert/go-crud?style=default&color=0080ff" alt="repo-language-count">

<!-- default option, no dependency badges. -->


<!-- default option, no dependency badges. -->

</div>
<br>

---

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
    - [Project Index](#project-index)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Testing](#testing)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

---

## Overview



---

## Features

|      | Component       | Details                              |
| :--- | :-------------- | :----------------------------------- |
| ⚙️  | **Architecture**  | <ul><li>The application follows a modular architecture, with main components like `main.go`, `handlers`, `models`, and `utils`.</li><li>It uses Go modules for dependency management, ensuring clear separation of dependencies in the `go.mod` and `go.sum` files.</li></ul> |
| 🔩 | **Code Quality**  | <ul><li>The codebase is well-organized with a clear directory structure, making it easy to navigate.</li><li>It adheres to Go's idiomatic practices, ensuring readability and maintainability.</li><li>Uses of `go fmt` and `golint` during development can help maintain consistent coding standards.</li></ul> |
| 📄 | **Documentation** | <ul><li>The project includes a basic README file with installation instructions and usage guidelines.</li><li>Code comments are present where necessary to explain complex functions or sections of code.</li></ul> |
| 🔌 | **Integrations**  | <ul><li>The application integrates with MySQL for database operations, utilizing the `database/sql` package along with a driver like `mysql`.</li><li>Uses `mux` (a lightweight HTTP router which matches incoming requests to their respective handler functions) for routing.</li></ul> |
| 🧩 | **Modularity**    | <ul><li>The codebase is highly modular, with each feature or functionality encapsulated in its own module or package.</li><li>For example, CRUD operations are handled in the `handlers` package, while database interactions are managed by the `models` package.</li></ul> |
| 📦 | **Dependencies**  | <ul><li>The project relies primarily on Go's standard library and community packages like `mux` for database operations (`mysql`), ensuring a lightweight dependency footprint.</li><li>Versioning dependencies is handled through Go modules, maintaining stability across different environments.</li></ul> |

---

## Project Structure

```sh
└── go-crud/
    ├── go.mod
    ├── go.sum
    ├── handlers.go
    ├── main.go
    ├── templates
    │   ├── addTaskForm.html
    │   ├── base.html
    │   ├── todoList.html
    │   └── updateTaskForm.html
    └── utility.go
```

### Project Index

<details open>
	<summary><b><code>GO-CRUD/</code></b></summary>
	<!-- __root__ Submodule -->
	<details>
		<summary><b>__root__</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ __root__</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/handlers.go'>handlers.go</a></b></td>
					<td style='padding: 8px;'>- The <code>handlers.go</code> file is pivotal to the web applications functionality by defining HTTP handlers that manage CRUD operations on tasks, templating with Gorilla Mux routers, and interacting with a database<br>- It facilitates task retrieval, creation, updating, and deletion through specific endpoints like <code>/fragment/{name}</code> for dynamic content loading and form submissions for adding or modifying tasks<br>- This file ensures seamless user interactions by handling requests and responses effectively, enhancing the applications responsiveness and data management capabilities.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/utility.go'>utility.go</a></b></td>
					<td style='padding: 8px;'>- GetTasksRetrieves all tasks from the database and returns them as a slice of Task structs<br>- This function queries the tasks" table and populates a list of tasks based on the results.2<br>- **GetTaskByIdFetches a single task by its ID<br>- It constructs a query to select the task with the specified ID and returns it as a pointer to a Task struct<br>- If no task is found, it returns an error indicating that the task does not exist.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/go.sum'>go.sum</a></b></td>
					<td style='padding: 8px;'>- The <code>go.sum</code> file serves as a crucial component of the projects dependency management system by recording exact versions of dependencies used during the build process<br>- This ensures that all team members and CI/CD pipelines use consistent and compatible package versions, maintaining consistency across different environments.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/go.mod'>go.mod</a></b></td>
					<td style='padding: 8px;'>- The <code>go.mod</code> file defines the module path and required dependencies for the Go project located at <code>github.com/plumpalbert/go-crud</code><br>- It specifies that the project uses libraries from <code>filippo.io</code>, <code>github.com/go-sql-driver/mysql</code>, and <code>github.com/gorilla/mux</code> to facilitate database operations and web server functionalities, respectively<br>- This setup ensures a robust foundation for building CRUD (Create, Read, Update, Delete) functionality in the application.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/main.go'>main.go</a></b></td>
					<td style='padding: 8px;'>- The <code>main.go</code> file serves as the entry point of the application, setting up a web server using Gorilla Mux for routing and handling HTTP requests<br>- It initializes a MySQL database connection, defines data structures (Task), and sets up handlers for CRUD operations on tasks via HTML templates rendered over HTTP<br>- This setup forms the backbone for an API/web service facilitating task management through a RESTful interface.</td>
				</tr>
			</table>
		</blockquote>
	</details>
	<!-- templates Submodule -->
	<details>
		<summary><b>templates</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ templates</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/templates/base.html'>base.html</a></b></td>
					<td style='padding: 8px;'>- The <code>templates/base.html</code> file defines the base HTML structure of an application that integrates HTMX and Tailwind CSS for a ToDo list, featuring dynamic task addition and display functionalities<br>- It includes a modal for adding new tasks, styled with Tailwind CSS, and utilizes HTMX to handle AJAX requests for updating the task list without reloading the page<br>- The script manages the visibility state of the modal by toggling its hidden class on button click events.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/templates/addTaskForm.html'>addTaskForm.html</a></b></td>
					<td style='padding: 8px;'>- The <code>addTaskForm.html</code> file defines a template for an add task form within the projects user interface<br>- This template provides a user interface element allowing users to input and submit new tasks, with options to cancel or confirm the addition of the task<br>- It is part of the frontend design, enhancing user interaction and data entry capabilities within the application.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/templates/todoList.html'>todoList.html</a></b></td>
					<td style='padding: 8px;'>- The <code>todoList.html</code> file defines a template that generates an HTML list of tasks from a data structure passed to it, including task names and actions like updating or deleting each task<br>- This template is part of a larger project aiming to create a web-based todo list application, facilitating the display and management of tasks through interactive buttons for updates and deletions.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/PlumpAlbert/go-crud/blob/master/templates/updateTaskForm.html'>updateTaskForm.html</a></b></td>
					<td style='padding: 8px;'>- The <code>updateTaskForm.html</code> file defines a template for an update task form within the projects user interface<br>- This template is designed to be used in a web application context, likely part of a larger system managing tasks or projects<br>- It includes fields for entering and updating task details such as name and completion status, with interactive elements like checkboxes and buttons facilitating data submission and interaction<br>- The file leverages HTML along with Go templating syntax (as indicated by the use of <code>{{.Task }}</code> and similar placeholders) to dynamically populate form fields based on existing data.</td>
				</tr>
			</table>
		</blockquote>
	</details>
</details>

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** Golang, HTML
- **Package Manager:** Go modules

### Installation

Build go-crud from the source and intsall dependencies:

1. **Clone the repository:**

    ```sh
    git clone https://github.com/PlumpAlbert/go-crud
    ```

2. **Navigate to the project directory:**

    ```sh
    cd go-crud
    ```

3. **Install the dependencies:**

	```sh
	go install
	```

### Usage

Run the project with:

```sh
go build ./cmd/gocrud
./gocrud
```

---

<div align="right">

[![][back-to-top]](#top)

</div>


[back-to-top]: https://img.shields.io/badge/-BACK_TO_TOP-151515?style=flat-square


---
