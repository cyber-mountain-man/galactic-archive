# 🌌 Galactic Archive

> *Preserving the knowledge of the cosmos — one entry at a time.*

**Galactic Archive** is a sci-fi inspired web application built with **Go**, **htmx**, and **Bulma**. Designed to feel like a futuristic knowledge interface from a galactic empire, the archive catalogs fictional planets, technologies, events, and species — paying tribute to classics like **Dune**, **Star Trek**, and the **Foundation series**.

This project is both a creative worldbuilding sandbox and a professional demonstration of modern full-stack web development using Go.

---

## 🚀 Stack Overview

| Layer        | Technology     | Purpose                                |
|--------------|----------------|----------------------------------------|
| Backend      | Go (`net/http`) | Server-side routing, templating, and logic |
| Frontend     | htmx           | Dynamic, reactive frontend without JS frameworks |
| CSS Framework| Bulma          | Clean, responsive UI with sci-fi vibes |
| Templates    | `html/template`| Componentized, server-rendered views   |
| Data Layer   | JSON / SQLite  | Portable storage of archive entries    |

---

## 🧪 Features

- 🪐 View and explore planets, factions, artifacts, and historical events
- 📜 Each entry is dynamically loaded using htmx (`hx-get`, `hx-swap`)
- 🧬 Search and filter entries by type, region, or status
- 🛠️ Admin panel (in progress): add and edit entries through the UI
- 🎨 Stylish retro-futuristic layout using Bulma and minimal custom CSS

---

## 💡 Learning Goals

This project was created to:

- Learn and apply the **Go + htmx** stack effectively
- Build a modular, extensible codebase for server-rendered web apps
- Explore creative frontend design using **Bulma** as a base
- Demonstrate real-world usage of `html/template`, dynamic routing, and lightweight client-server interactions

---

## 📂 Project Structure

```
galactic-archive/
├── main.go                  # Main application file
├── templates/               # HTML templates rendered by Go
│   ├── layout.html
│   ├── index.html
│   └── entry.html
├── static/                  # Static files (CSS, JS)
│   ├── css/
│   │   └── bulma.min.css
│   └── js/
│       └── htmx.min.js
├── data/                    # Mock data in JSON or SQLite
└── README.md
```

---

## 🧠 Inspiration

This app is inspired by the idea of a **central galactic database** — like the Federation's Memory Alpha, the Imperial Archives of Trantor, or the Bene Gesserit’s hidden knowledge repositories. It’s a tribute to speculative fiction’s belief in the preservation of knowledge and the beauty of structured information.

---

## 🛠 Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/<your-username>/galactic-archive.git
   cd galactic-archive
   ```

2. **Run the server**
   ```bash
   go run main.go
   ```

3. **Visit** `http://localhost:8080` in your browser and begin your journey through the stars ✨

---

## 🧩 Future Features

- 🛡️ Entry classification (e.g. Public, Classified, Redacted)
- 👤 Archivist authentication system
- 📡 Live feed of “galactic events”
- 🎙️ Audio or visual log entries (mocked or uploaded)
- 🌌 Dark mode and theming support

---

## 👨‍🚀 Author

**Guillermo Morrison**  
Cybersecurity & Software Development  
[GitHub @CyberMountainMan](https://github.com/CyberMountainMan)

---

## 📄 License

This project is licensed under the **MIT License** — feel free to fork, remix, and build your own galactic records.

---

> _"The true strength of a civilization is not measured by its weapons or its armies... but by the knowledge it preserves for future generations."_  
> — *Archivist's Creed, Circa 3412 A.G.*

