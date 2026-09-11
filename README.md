<p align="center">
  <img src="web/static/aura-icon.svg" width="88" height="88" alt="AURA">
</p>

<h1 align="center">AURA</h1>

<p align="center">
  <strong>Advanced URL and Request Analyzer</strong><br>
  Продвинутый анализатор URL и запросов
</p>

<p align="center">
  Локальный HTTP/1.1 прокси, карта цели и рабочий стол для своей лаборатории.
</p>

<p align="center">
  <a href="#en">English</a> ·
  <a href="#ru">Русский</a> ·
  <a href="ARCHITECTURE.md">Architecture</a> ·
  <a href="LICENSE">License</a>
</p>

---

<a id="en"></a>
## English

AURA is a **desktop GTK/WebKit app** plus a local API. Point a browser at `127.0.0.1:8080`, confirm the host is yours, and build a **Site map** from proxy history and recon modules.

| Listen | Default | What it is |
| --- | --- | --- |
| UI / API | `127.0.0.1:1337` | Map, Proxy, Replay, Payloads… |
| Proxy | `127.0.0.1:8080` | HTTP/1.1 intercept (no HTTP/2 MITM) |
| Callback | `127.0.0.1:8082` | Out-of-band hits for your own tests |

Active probes (DNS brute, ports, dirs, params) stay off until you check **This is my target / lab**.

### Requirements

- Go 1.22+
- Linux with GTK 3 and WebKitGTK 4.1 (Kali/Debian)

```bash
sudo apt install golang-go gcc pkg-config libgtk-3-dev libwebkit2gtk-4.1-dev
git clone git@github.com:msdrakula/AURA.git
cd AURA
CGO_ENABLED=1 go build -o aura ./cmd/server
./aura
```

Menu entry on Kali:

```bash
sh packaging/kali/install-command.sh
```

### HTTPS

1. Open **Proxy → Settings**, download `aura-ca.crt`.
2. Trust that CA in the browser / OS.
3. Set HTTP proxy to `127.0.0.1:8080`.

A fresh CA is created on first run under `data/`. Do not copy another machine’s CA keys.

### Map

1. Browse through the proxy — hosts and paths land on **Map → Site map**.
2. Type the host, check the lab box, click **Start**.
3. Run modules (subdomains, DNS, live hosts, ports, tech, Wayback, scrape, dirs, params). Findings merge into the same tree.
4. **Scope** is a separate tab and filters the tree.

### Wordlists

Optional [SecLists](https://github.com/danielmiessler/SecLists) (MIT):

```bash
git clone --depth 1 https://github.com/danielmiessler/SecLists.git third_party/SecLists
```

Discover / Fuzz / Map pick files from disk. Built-in short lists work without SecLists.

### Flags

```bash
./aura \
  --api-port 1337 \
  --proxy-host 127.0.0.1 \
  --proxy-port 8080 \
  --callback-port 8082 \
  --db-path aura.db \
  --ca-dir data \
  --window=true
```

`--window=false` serves the UI in a normal browser tab.

### License

[PolyForm Noncommercial 1.0.0](LICENSE) — personal / noncommercial use. Companies and paid work need written permission. Plain language: [LICENSE-NOTES.md](LICENSE-NOTES.md).

---

<a id="ru"></a>
## Русский

AURA — это **окно GTK** и локальный прокси. Трафик браузера идёт через вас. Карта цели собирается из истории прокси и модулей разведки.

Используйте только **свой** трафик и лабораторные стенды.

### Вкладки

| Вкладка | Зачем |
| --- | --- |
| **Map** | Site map + модули (поддомены, DNS, порты, tech, Wayback, HTML/JS, пути, параметры) |
| **Scope** | Включить / исключить хосты |
| **Proxy** | Перехват, HTTP history, match & replace |
| **Discover / Fuzz** | Перебор путей и подстановка `FUZZ` |
| **Payloads / Replay** | Пакетная подстановка и один запрос туда-обратно |
| **Scanner** | Пассивный разбор запроса |
| **Decoder / Diff / Tokens** | Кодировки, сравнение, энтропия |
| **Callback** | Локальный колбэк `8082` |
| **Saved / Other** | Сохранённые запросы и находки analyzer |

### Сборка

```bash
sudo apt install golang-go gcc pkg-config libgtk-3-dev libwebkit2gtk-4.1-dev
git clone git@github.com:msdrakula/AURA.git
cd AURA
CGO_ENABLED=1 go build -o aura ./cmd/server
./aura
```

- Интерфейс: http://127.0.0.1:1337
- Прокси: `127.0.0.1:8080`
- Колбэк: `127.0.0.1:8082`

Старый `meb.db` подхватится сам, пока нет `aura.db`.

### Карта цели

1. Ходите сайтом через прокси — дерево на **Map → Site map** заполняется сразу.
2. Вставьте хост, отметьте «это моя цель / лаборатория», нажмите **Начать**.
3. Откройте модуль, настройте timeout / wordlist / порты, **Запустить**.
4. Находки модулей дописываются в то же дерево, что и история прокси.

Пассивные шаги ходят в открытые источники (crt.sh, Wayback). Активные шлют запросы только на подтверждённую цель.

### HTTPS

Скачайте CA со вкладки Proxy (`aura-ca.crt`), добавьте в доверенные, укажите прокси `127.0.0.1:8080`. Если раньше стоял старый CA — импортируйте новый **AURA Intercept CA**.

---

## Репозиторий

В git — исходники, UI, тесты, упаковка под Kali.

На каждой машине **заново** появляются (и не хранятся в git):

- `data/ca/`, `data/leaf/` — локальный MITM CA
- `aura.db` / `meb.db` — история прокси
- `data/debug/*.jsonl` — журнал сессии
- бинарник `aura` после `go build`
- `third_party/SecLists/` — клонируется отдельно

Это не «секрет продукта», это **ваши** ключи и трафик с этого компьютера. Дома AURA создаст свои.

Прокси работает по **HTTP/1.1**.

### Лицензия

[PolyForm Noncommercial 1.0.0](LICENSE) — для себя и некоммерции. Компаниям и платной работе нужно письменное разрешение. Простыми словами: [LICENSE-NOTES.md](LICENSE-NOTES.md).
