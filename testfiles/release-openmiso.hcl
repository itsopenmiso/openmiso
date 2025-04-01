project = "openmiso"

app "openmiso" {

  build {
    use "exec" {}
  }

  release {
    use "github" {
      files = [
      "./openmiso",
      "./openmiso-server",
      "./openmiso-runner",
      ]
    }
  }
}
