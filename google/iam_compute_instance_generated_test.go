// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/compute.osLogin",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/compute.osLogin",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/compute.osLogin",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_member" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_instance_iam_policy" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_compute_instance_iam_policy" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_binding" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_binding" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
