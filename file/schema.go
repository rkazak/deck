// Code generated by go generate; DO NOT EDIT.
package file

const contentSchema = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "_format_version": {
      "type": "string"
    },
    "_info": {
      "$schema": "http://json-schema.org/draft-04/schema#",
      "$ref": "#/definitions/Info"
    },
    "_plugin_configs": {
      "patternProperties": {
        ".*": {
          "additionalProperties": true,
          "type": "object"
        }
      },
      "type": "object"
    },
    "_workspace": {
      "type": "string"
    },
    "ca_certificates": {
      "items": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "$ref": "#/definitions/FCACertificate"
      },
      "type": "array"
    },
    "certificates": {
      "items": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "$ref": "#/definitions/FCertificate"
      },
      "type": "array"
    },
    "consumers": {
      "items": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "$ref": "#/definitions/FConsumer"
      },
      "type": "array"
    },
    "plugins": {
      "items": {
        "$ref": "#/definitions/FPlugin"
      },
      "type": "array"
    },
    "routes": {
      "items": {
        "$ref": "#/definitions/FRoute"
      },
      "type": "array"
    },
    "services": {
      "items": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "$ref": "#/definitions/FService"
      },
      "type": "array"
    },
    "upstreams": {
      "items": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "$ref": "#/definitions/FUpstream"
      },
      "type": "array"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "definitions": {
    "ACLGroup": {
      "required": [
        "group"
      ],
      "properties": {
        "consumer": {
          "$ref": "#/definitions/Consumer"
        },
        "created_at": {
          "type": "integer"
        },
        "group": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "ActiveHealthcheck": {
      "properties": {
        "concurrency": {
          "type": "integer"
        },
        "healthy": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/Healthy"
        },
        "http_path": {
          "type": "string"
        },
        "https_sni": {
          "type": "string"
        },
        "https_verify_certificate": {
          "type": "boolean"
        },
        "timeout": {
          "type": "integer"
        },
        "type": {
          "type": "string"
        },
        "unhealthy": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/Unhealthy"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "BasicAuth": {
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "consumer": {
          "$ref": "#/definitions/Consumer"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "username": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "CIDRPort": {
      "properties": {
        "ip": {
          "type": "string"
        },
        "port": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Certificate": {
      "properties": {
        "cert": {
          "type": "string"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "key": {
          "type": "string"
        },
        "snis": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Consumer": {
      "properties": {
        "created_at": {
          "type": "integer"
        },
        "custom_id": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "username": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "anyOf": [
        {
          "required": [
            "id"
          ]
        },
        {
          "required": [
            "username"
          ]
        }
      ]
    },
    "FCACertificate": {
      "required": [
        "cert"
      ],
      "properties": {
        "cert": {
          "type": "string"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FCertificate": {
      "required": [
        "id",
        "cert",
        "key"
      ],
      "properties": {
        "cert": {
          "type": "string"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "key": {
          "type": "string"
        },
        "snis": {
          "items": {
            "properties": {
              "name": {
                "type": "string"
              }
            },
            "type": "object"
          },
          "type": "array"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FConsumer": {
      "properties": {
        "acls": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/ACLGroup"
          },
          "type": "array"
        },
        "basicauth_credentials": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/BasicAuth"
          },
          "type": "array"
        },
        "created_at": {
          "type": "integer"
        },
        "custom_id": {
          "type": "string"
        },
        "hmacauth_credentials": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/HMACAuth"
          },
          "type": "array"
        },
        "id": {
          "type": "string"
        },
        "jwt_secrets": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/JWTAuth"
          },
          "type": "array"
        },
        "keyauth_credentials": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/KeyAuth"
          },
          "type": "array"
        },
        "oauth2_credentials": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/Oauth2Credential"
          },
          "type": "array"
        },
        "plugins": {
          "items": {
            "$ref": "#/definitions/FPlugin"
          },
          "type": "array"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "username": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FPlugin": {
      "required": [
        "name"
      ],
      "properties": {
        "_config": {
          "type": "string"
        },
        "config": {
          "additionalProperties": true,
          "type": "object"
        },
        "consumer": {
          "type": "string"
        },
        "created_at": {
          "type": "integer"
        },
        "enabled": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "protocols": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "route": {
          "type": "string"
        },
        "run_on": {
          "type": "string"
        },
        "service": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FRoute": {
      "properties": {
        "created_at": {
          "type": "integer"
        },
        "destinations": {
          "items": {
            "$ref": "#/definitions/CIDRPort"
          },
          "type": "array"
        },
        "headers": {
          "patternProperties": {
            ".*": {
              "items": {
                "type": "string"
              },
              "type": "array"
            }
          },
          "type": "object"
        },
        "hosts": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "https_redirect_status_code": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "methods": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "name": {
          "type": "string"
        },
        "path_handling": {
          "type": "string"
        },
        "paths": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "plugins": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/FPlugin"
          },
          "type": "array"
        },
        "preserve_host": {
          "type": "boolean"
        },
        "protocols": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "regex_priority": {
          "type": "integer"
        },
        "service": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/Service"
        },
        "snis": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "sources": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/CIDRPort"
          },
          "type": "array"
        },
        "strip_path": {
          "type": "boolean"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "updated_at": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FService": {
      "properties": {
        "client_certificate": {
          "type": "string"
        },
        "connect_timeout": {
          "type": "integer"
        },
        "created_at": {
          "type": "integer"
        },
        "host": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "plugins": {
          "items": {
            "$ref": "#/definitions/FPlugin"
          },
          "type": "array"
        },
        "port": {
          "type": "integer"
        },
        "protocol": {
          "type": "string"
        },
        "read_timeout": {
          "type": "integer"
        },
        "retries": {
          "type": "integer"
        },
        "routes": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/FRoute"
          },
          "type": "array"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "updated_at": {
          "type": "integer"
        },
        "url": {
          "type": "string"
        },
        "write_timeout": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FTarget": {
      "required": [
        "target"
      ],
      "properties": {
        "created_at": {
          "type": "number"
        },
        "id": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "target": {
          "type": "string"
        },
        "upstream": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/Upstream"
        },
        "weight": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "FUpstream": {
      "properties": {
        "algorithm": {
          "type": "string"
        },
        "created_at": {
          "type": "integer"
        },
        "hash_fallback": {
          "type": "string"
        },
        "hash_fallback_header": {
          "type": "string"
        },
        "hash_on": {
          "type": "string"
        },
        "hash_on_cookie": {
          "type": "string"
        },
        "hash_on_cookie_path": {
          "type": "string"
        },
        "hash_on_header": {
          "type": "string"
        },
        "healthchecks": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/Healthcheck"
        },
        "host_header": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "slots": {
          "type": "integer"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "targets": {
          "items": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "$ref": "#/definitions/FTarget"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "HMACAuth": {
      "required": [
        "username",
        "secret"
      ],
      "properties": {
        "consumer": {
          "$ref": "#/definitions/Consumer"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "username": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Healthcheck": {
      "properties": {
        "active": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/ActiveHealthcheck"
        },
        "passive": {
          "$schema": "http://json-schema.org/draft-04/schema#",
          "$ref": "#/definitions/PassiveHealthcheck"
        },
        "threshold": {
          "type": "number"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Healthy": {
      "properties": {
        "http_statuses": {
          "items": {
            "type": "integer"
          },
          "type": "array"
        },
        "interval": {
          "type": "integer"
        },
        "successes": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Info": {
      "properties": {
        "select_tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "JWTAuth": {
      "required": [
        "algorithm",
        "key",
        "secret"
      ],
      "properties": {
        "algorithm": {
          "type": "string"
        },
        "consumer": {
          "$ref": "#/definitions/Consumer"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "key": {
          "type": "string"
        },
        "rsa_public_key": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "KeyAuth": {
      "required": [
        "key"
      ],
      "properties": {
        "consumer": {
          "$ref": "#/definitions/Consumer"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "key": {
          "type": "string"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Oauth2Credential": {
      "required": [
        "name",
        "client_id",
        "redirect_uris",
        "client_secret"
      ],
      "properties": {
        "client_id": {
          "type": "string"
        },
        "client_secret": {
          "type": "string"
        },
        "consumer": {
          "$ref": "#/definitions/Consumer"
        },
        "created_at": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "redirect_uris": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "PassiveHealthcheck": {
      "properties": {
        "healthy": {
          "$ref": "#/definitions/Healthy"
        },
        "unhealthy": {
          "$ref": "#/definitions/Unhealthy"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Route": {
      "properties": {
        "created_at": {
          "type": "integer"
        },
        "destinations": {
          "items": {
            "$ref": "#/definitions/CIDRPort"
          },
          "type": "array"
        },
        "headers": {
          "patternProperties": {
            ".*": {
              "items": {
                "type": "string"
              },
              "type": "array"
            }
          },
          "type": "object"
        },
        "hosts": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "https_redirect_status_code": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        },
        "methods": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "name": {
          "type": "string"
        },
        "path_handling": {
          "type": "string"
        },
        "paths": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "preserve_host": {
          "type": "boolean"
        },
        "protocols": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "regex_priority": {
          "type": "integer"
        },
        "service": {
          "$ref": "#/definitions/Service"
        },
        "snis": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "sources": {
          "items": {
            "$ref": "#/definitions/CIDRPort"
          },
          "type": "array"
        },
        "strip_path": {
          "type": "boolean"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "updated_at": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "anyOf": [
        {
          "required": [
            "id"
          ]
        },
        {
          "required": [
            "name"
          ]
        }
      ]
    },
    "Service": {
      "properties": {
        "client_certificate": {
          "$ref": "#/definitions/Certificate"
        },
        "connect_timeout": {
          "type": "integer"
        },
        "created_at": {
          "type": "integer"
        },
        "host": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "port": {
          "type": "integer"
        },
        "protocol": {
          "type": "string"
        },
        "read_timeout": {
          "type": "integer"
        },
        "retries": {
          "type": "integer"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "updated_at": {
          "type": "integer"
        },
        "write_timeout": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "anyOf": [
        {
          "required": [
            "id"
          ]
        },
        {
          "required": [
            "name"
          ]
        }
      ]
    },
    "Unhealthy": {
      "properties": {
        "http_failures": {
          "type": "integer"
        },
        "http_statuses": {
          "items": {
            "type": "integer"
          },
          "type": "array"
        },
        "interval": {
          "type": "integer"
        },
        "tcp_failures": {
          "type": "integer"
        },
        "timeouts": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Upstream": {
      "required": [
        "name"
      ],
      "properties": {
        "algorithm": {
          "type": "string"
        },
        "created_at": {
          "type": "integer"
        },
        "hash_fallback": {
          "type": "string"
        },
        "hash_fallback_header": {
          "type": "string"
        },
        "hash_on": {
          "type": "string"
        },
        "hash_on_cookie": {
          "type": "string"
        },
        "hash_on_cookie_path": {
          "type": "string"
        },
        "hash_on_header": {
          "type": "string"
        },
        "healthchecks": {
          "$ref": "#/definitions/Healthcheck"
        },
        "host_header": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "slots": {
          "type": "integer"
        },
        "tags": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object"
    }
  }
}`