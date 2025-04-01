# 🥣 openMiso

**openMiso** is a simple, open-source platform for building, deploying, and
running your software — inspired
by [HashiCorp Waypoint](https://github.com/hashicorp/waypoint), reimagined with
taste.

> **Why openMiso?**  
> Because good systems, like good soup, start with a great base.

---

## 🍜 What is openMiso?

openMiso is a minimal, modular alternative to heavy CI/CD systems and PaaS-like
platforms. It provides a **developer-first experience** for describing and
running your build and deploy workflows, powered by a **plugin-based
architecture** and an elegant, evolving core.

openMiso is:

- 🧩 **Plugin-first** – everything is a plugin: builders, deployers, runners
- 📦 **Packaged for portability** – local-first, server-optional
- 🔁 **Reproducible** – build once, deploy anywhere
- 📡 **Optional server mode** – run openMiso in CLI or server form
- 🌱 **PoC-friendly** – start with JSON over HTTP, scale up to WebSockets and
  gRPC

---

## 💡 Why "openMiso"?

- **Miso paste** is the essential base for many soups in Japanese cuisine.
- Like miso in the kitchen, `miso` in your pipeline is a **fermented foundation
  ** — simple, flexible, and rich in umami.
- It's the thing you can build on top of, **again and again**, regardless of
  what's in your bowl (or repo).

Just as:
> miso + tofu + wakame = miso soup
> miso + docker + ssh = deployment
> **Miso** is the base. **You** add the flavor.

---

## 🚀 Status

🧪 **Proof of concept** in progress.  
Currently building:

- `miso-cli` – CLI tool
- `miso-server` – simple HTTP server
- `miso-plugin-sdk` – plugin contract & types
- `miso-plugin-hello` – test plugin
- JSON-based plugin protocol via `stdin/stdout`
- WebSocket logging support (runner ↔ server)
- gRPC planned (but not yet)

---

## 📦 Core Repositories (planned or created)

| Name                                                              | Description       |
|-------------------------------------------------------------------|-------------------|
| [`itsopenmiso/openmiso`](https://github.com/itsopenmiso/openmiso) | CLI + Server core |

---

## 📐 Roadmap

- [x] JSON over HTTP for core communication
- [x] Plugin execution via subprocess
- [ ] WebSocket support for streaming logs
- [ ] Runner agent with reconnecting capabilities
- [ ] gRPC internal protocol (server ↔ runner)
- [ ] GUI (Textual or web-based)
- [ ] Deployment registry / state store
- [ ] OpenMiso Cloud™ (maybe someday 😄)

---

## 🤝 Inspired by

- [Waypoint](https://github.com/hashicorp/waypoint) by HashiCorp
- [Nomad](https://github.com/hashicorp/nomad)
- [Dagger](https://github.com/dagger/dagger)
- The gentle beauty of miso soup

---

## 🧠 Philosophy

We believe developer tooling should be:

- **understandable** without reading a 400-page manual
- **composable**, not monolithic
- **debuggable** without Wireshark
- **portable**, not tied to vendor clouds

openMiso isn't here to replace everything.  
It's here to **make the small things beautiful again**.

---

## 🥢 Contributing

We're still fermenting the early stages of this project.  
Pull requests, ideas, comments, plugins, memes — all welcome.

> 🍥 Let’s cook something great, one spoon at a time.

---

## 📞 Contact

Want to collaborate?  
Say hi at `github.com/itsopenmiso/openmiso` or open an issue!

---

## Licensing

openMiso is licensed under the Mozilla Public License 2.0 (MPL-2.0).

We chose MPL because it ensures modifications to this codebase stay open,  
while allowing adoption in a wide range of projects and environments.

We reserve the right to relicense future versions under a stronger copyleft
license (such as AGPLv3)  
to protect the values of openness and shared ownership that define Miso.

*
