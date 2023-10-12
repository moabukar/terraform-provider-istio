provider istio {
  config_path = "/Users/mohameda/.kube/config"
}

resource "istio_virtual_service" "test"{
  metadata {
    name = "istio-example"
    namespace = "test"
  }

  spec {
    hosts = [
      "test.prod.svc.cluster.local",
    ]
    http  {
        name = "test-route"
        route  {
            destination {
              host = "test.prod.svc.cluster.local"
              subset = "v1"
            }
          }
       }
    }
}
