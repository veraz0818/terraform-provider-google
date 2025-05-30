// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package privateca_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccPrivatecaCaPool_privatecaCapoolBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCaPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPool_privatecaCapoolBasicExample(context),
			},
			{
				ResourceName:            "google_privateca_ca_pool.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCaPool_privatecaCapoolBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}
`, context)
}

func TestAccPrivatecaCaPool_privatecaCapoolAllFieldsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCaPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPool_privatecaCapoolAllFieldsExample(context),
			},
			{
				ResourceName:            "google_privateca_ca_pool.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCaPool_privatecaCapoolAllFieldsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = false
    publish_crl = true
    encoding_format = "PEM"
  }
  labels = {
    foo = "bar"
  }
  issuance_policy {
    allowed_key_types {
      elliptic_curve {
        signature_algorithm = "ECDSA_P256"
      }
    }
    allowed_key_types {
      rsa {
        min_modulus_size = 5
        max_modulus_size = 10
      }
    }
    backdate_duration = "3600s"
    maximum_lifetime = "50000s"
    allowed_issuance_modes {
      allow_csr_based_issuance = true
      allow_config_based_issuance = true
    }
    identity_constraints {
      allow_subject_passthrough = true
      allow_subject_alt_names_passthrough = true
      cel_expression {
        expression = "subject_alt_names.all(san, san.type == DNS || san.type == EMAIL )"
        title = "My title"
      }
    }
    baseline_values {
      aia_ocsp_servers = ["example.com"]
      additional_extensions {
        critical = true
        value = "asdf"
        object_id {
          object_id_path = [1, 7]
        }
      }
      policy_ids {
        object_id_path = [1, 5]
      }
      policy_ids {
        object_id_path = [1, 5, 7]
      }
      ca_options {
        is_ca = true
        max_issuer_path_length = 10
      }
      key_usage {
        base_key_usage {
          digital_signature = true
          content_commitment = true
          key_encipherment = false
          data_encipherment = true
          key_agreement = true
          cert_sign = false
          crl_sign = true
          decipher_only = true
        }
        extended_key_usage {
          server_auth = true
          client_auth = false
          email_protection = true
          code_signing = true
          time_stamping = true
        }
      }
      name_constraints {
        critical                  = true
        permitted_dns_names       = ["*.example1.com", "*.example2.com"]
        excluded_dns_names        = ["*.deny.example1.com", "*.deny.example2.com"]
        permitted_ip_ranges       = ["10.0.0.0/8", "11.0.0.0/8"]
        excluded_ip_ranges        = ["10.1.1.0/24", "11.1.1.0/24"]
        permitted_email_addresses = [".example1.com", ".example2.com"]
        excluded_email_addresses  = [".deny.example1.com", ".deny.example2.com"]
        permitted_uris            = [".example1.com", ".example2.com"]
        excluded_uris             = [".deny.example1.com", ".deny.example2.com"]
      }
    }
  }
}
`, context)
}

func testAccCheckPrivatecaCaPoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_privateca_ca_pool" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PrivatecaBasePath}}projects/{{project}}/locations/{{location}}/caPools/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("PrivatecaCaPool still exists at %s", url)
			}
		}

		return nil
	}
}
