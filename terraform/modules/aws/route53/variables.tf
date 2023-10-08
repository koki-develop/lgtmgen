variable "domains" {
  type = object({
    apex   = string
    ui     = optional(string, null)
    api    = string
    images = string
  })
}

variable "routings" {
  type = object({
    api = object({
      domain_name = string
      zone_id     = string
    })
    images = object({
      domain_name = string
      zone_id     = string
    })
  })
}
