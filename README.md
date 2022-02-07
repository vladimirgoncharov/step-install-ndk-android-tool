# Install Android NDK component

[![Step changelog](https://shields.io/github/v/release/vladimirgoncharov/step-install-ndk-android-tool?include_prereleases&label=changelog&color=blueviolet)](https://github.com/vladimirgoncharov/step-install-ndk-android-tool/releases)

Install Android NDK component that are required for the app.

<details>
<summary>Description</summary>


### Configuring the Step

1. Set its revision in the **NDK Revision** input.

### Troubleshooting

If the Step fails, check that your repo actually contains a `gradlew` file. Without the Gradle wrapper, this Step won't work.
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `ndk_version` | NDK version to install, for example `23.1.7779620`. Run `sdkmanager --list` on your machine to see all available versions. Leave this input empty if you are not using the Native Development Kit in your project. | required | `23.1.7779620` |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/vladimirgoncharov/step-install-ndk-android-tool/pulls) and [issues](https://github.com/vladimirgoncharov/step-install-ndk-android-tool/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
