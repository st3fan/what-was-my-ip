job "what-was-my-ip-lookup" {
  datacenters = ["whitby"]

  type = "batch"

  periodic {
    crons = ["0 */15 * * * * *"]
    prohibit_overlap = true
  }

  group "what-was-my-ip-lookup" {
    task "what-was-my-ip-lookup" {
      driver = "docker"

      resources {
        memory = 64
      }

      vault {
        policies = ["access-secrets"]
      }

      config {
        image = "ghcr.io/st3fan/what-was-my-ip:main"
        args = ["lookup"]
      }

      template {
        data = <<EOF
         {{ with secret "secrets/what-was-my-ip/database" }}
           DB_USERNAME={{ .Data.data.username | toJSON }}
           DB_PASSWORD={{ .Data.data.password | toJSON }}
           DB_HOSTNAME={{ .Data.data.hostname | toJSON }}
           DB_DATABASE={{ .Data.data.database | toJSON }}
         {{ end }}
        EOF
        destination = "secrets/database.env"
        env = true
      }
    }
  }
}
