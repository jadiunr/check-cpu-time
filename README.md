[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/jadiunr/check-cpu-time)
![Go Test](https://github.com/jadiunr/check-cpu-time/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/jadiunr/check-cpu-time/workflows/goreleaser/badge.svg)

# Check Plugin Template

## Overview
check-plugin-template is a template repository which wraps the [Sensu Plugin SDK][2].
To use this project as a template, click the "Use this template" button from the main project page.
Once the repository is created from this template, you can use the [Sensu Plugin Tool][9] to
populate the templated fields with the proper values.

## Functionality

After successfully creating a project from this template, update the `Config` struct with any
configuration options for the plugin, map those values as plugin options in the variable `options`,
and customize the `checkArgs` and `executeCheck` functions in [main.go][7].

When writing or updating a plugin's README from this template, review the Sensu Community
[plugin README style guide][3] for content suggestions and guidance. Remove everything
prior to `# Sensu CPU time check` from the generated README file, and add additional context about the
plugin per the style guide.

## Releases with Github Actions

To release a version of your project, simply tag the target sha with a semver release without a `v`
prefix (ex. `1.0.0`). This will trigger the [GitHub action][5] workflow to [build and release][4]
the plugin with goreleaser. Register the asset with [Bonsai][8] to share it with the community!

***

# Sensu CPU time check

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The Sensu CPU time check is a [Sensu Check][6] that ...

## Files

## Usage examples

<details>
<summary>Check output</summary>

```
# HELP cpu_time_total Static CPU Total time
# TYPE cpu_time_total counter
cpu_time_total{cpu="cpu0"} 1215.95
cpu_time_total{cpu="cpu1"} 1217.53
cpu_time_total{cpu="cpu2"} 1217.36
cpu_time_total{cpu="cpu3"} 1217.09
cpu_time_total{cpu="cpu4"} 1217.91
cpu_time_total{cpu="cpu5"} 1218.79
cpu_time_total{cpu="cpu6"} 1218.37
cpu_time_total{cpu="cpu7"} 1218.75
cpu_time_total{cpu="cpu8"} 1219.60
cpu_time_total{cpu="cpu9"} 1218.68
cpu_time_total{cpu="cpu10"} 1210.61
cpu_time_total{cpu="cpu11"} 1215.81
cpu_time_total{cpu="cpu-total"} 14606.82
# HELP cpu_time_user Statistic CPU User time
# TYPE cpu_time_user counter
cpu_time_user{cpu="cpu0"} 91.82
cpu_time_user{cpu="cpu1"} 87.37
cpu_time_user{cpu="cpu2"} 94.92
cpu_time_user{cpu="cpu3"} 127.36
cpu_time_user{cpu="cpu4"} 106.02
cpu_time_user{cpu="cpu5"} 101.07
cpu_time_user{cpu="cpu6"} 107.33
cpu_time_user{cpu="cpu7"} 100.62
cpu_time_user{cpu="cpu8"} 108.71
cpu_time_user{cpu="cpu9"} 101.25
cpu_time_user{cpu="cpu10"} 65.64
cpu_time_user{cpu="cpu11"} 85.07
cpu_time_user{cpu="cpu-total"} 1177.23
# HELP cpu_time_system Statistic CPU System time
# TYPE cpu_time_system counter
cpu_time_system{cpu="cpu0"} 26.84
cpu_time_system{cpu="cpu1"} 25.59
cpu_time_system{cpu="cpu2"} 24.54
cpu_time_system{cpu="cpu3"} 24.81
cpu_time_system{cpu="cpu4"} 25.68
cpu_time_system{cpu="cpu5"} 24.85
cpu_time_system{cpu="cpu6"} 24.13
cpu_time_system{cpu="cpu7"} 23.77
cpu_time_system{cpu="cpu8"} 23.54
cpu_time_system{cpu="cpu9"} 23.76
cpu_time_system{cpu="cpu10"} 26.54
cpu_time_system{cpu="cpu11"} 28.59
cpu_time_system{cpu="cpu-total"} 302.71
# HELP cpu_time_idle Statistic CPU Idle time
# TYPE cpu_time_idle counter
cpu_time_idle{cpu="cpu0"} 1091.55
cpu_time_idle{cpu="cpu1"} 1097.39
cpu_time_idle{cpu="cpu2"} 1093.90
cpu_time_idle{cpu="cpu3"} 1060.68
cpu_time_idle{cpu="cpu4"} 1082.18
cpu_time_idle{cpu="cpu5"} 1088.98
cpu_time_idle{cpu="cpu6"} 1082.32
cpu_time_idle{cpu="cpu7"} 1090.73
cpu_time_idle{cpu="cpu8"} 1083.70
cpu_time_idle{cpu="cpu9"} 1089.82
cpu_time_idle{cpu="cpu10"} 1097.20
cpu_time_idle{cpu="cpu11"} 1094.44
cpu_time_idle{cpu="cpu-total"} 13052.93
# HELP cpu_time_nice Statistic CPU Nice time
# TYPE cpu_time_nice counter
cpu_time_nice{cpu="cpu0"} 0.00
cpu_time_nice{cpu="cpu1"} 0.00
cpu_time_nice{cpu="cpu2"} 0.00
cpu_time_nice{cpu="cpu3"} 0.00
cpu_time_nice{cpu="cpu4"} 0.00
cpu_time_nice{cpu="cpu5"} 0.00
cpu_time_nice{cpu="cpu6"} 0.00
cpu_time_nice{cpu="cpu7"} 0.00
cpu_time_nice{cpu="cpu8"} 0.00
cpu_time_nice{cpu="cpu9"} 0.00
cpu_time_nice{cpu="cpu10"} 0.00
cpu_time_nice{cpu="cpu11"} 0.01
cpu_time_nice{cpu="cpu-total"} 0.05
# HELP cpu_time_iowait Statistic CPU Iowait time
# TYPE cpu_time_iowait counter
cpu_time_iowait{cpu="cpu0"} 0.99
cpu_time_iowait{cpu="cpu1"} 0.85
cpu_time_iowait{cpu="cpu2"} 0.80
cpu_time_iowait{cpu="cpu3"} 1.04
cpu_time_iowait{cpu="cpu4"} 0.83
cpu_time_iowait{cpu="cpu5"} 0.81
cpu_time_iowait{cpu="cpu6"} 1.11
cpu_time_iowait{cpu="cpu7"} 0.82
cpu_time_iowait{cpu="cpu8"} 0.74
cpu_time_iowait{cpu="cpu9"} 0.97
cpu_time_iowait{cpu="cpu10"} 0.78
cpu_time_iowait{cpu="cpu11"} 1.13
cpu_time_iowait{cpu="cpu-total"} 10.93
# HELP cpu_time_irq Statistic CPU Irq time
# TYPE cpu_time_irq counter
cpu_time_irq{cpu="cpu0"} 2.32
cpu_time_irq{cpu="cpu1"} 4.75
cpu_time_irq{cpu="cpu2"} 1.72
cpu_time_irq{cpu="cpu3"} 1.78
cpu_time_irq{cpu="cpu4"} 1.77
cpu_time_irq{cpu="cpu5"} 1.71
cpu_time_irq{cpu="cpu6"} 2.00
cpu_time_irq{cpu="cpu7"} 1.60
cpu_time_irq{cpu="cpu8"} 1.64
cpu_time_irq{cpu="cpu9"} 1.68
cpu_time_irq{cpu="cpu10"} 15.83
cpu_time_irq{cpu="cpu11"} 2.74
cpu_time_irq{cpu="cpu-total"} 39.60
# HELP cpu_time_softirq Statistic CPU Softirq time
# TYPE cpu_time_softirq counter
cpu_time_softirq{cpu="cpu0"} 2.43
cpu_time_softirq{cpu="cpu1"} 1.58
cpu_time_softirq{cpu="cpu2"} 1.48
cpu_time_softirq{cpu="cpu3"} 1.42
cpu_time_softirq{cpu="cpu4"} 1.43
cpu_time_softirq{cpu="cpu5"} 1.37
cpu_time_softirq{cpu="cpu6"} 1.48
cpu_time_softirq{cpu="cpu7"} 1.21
cpu_time_softirq{cpu="cpu8"} 1.27
cpu_time_softirq{cpu="cpu9"} 1.20
cpu_time_softirq{cpu="cpu10"} 4.62
cpu_time_softirq{cpu="cpu11"} 3.83
cpu_time_softirq{cpu="cpu-total"} 23.37
# HELP cpu_time_steal Statistic CPU Steal time
# TYPE cpu_time_steal counter
cpu_time_steal{cpu="cpu0"} 0.00
cpu_time_steal{cpu="cpu1"} 0.00
cpu_time_steal{cpu="cpu2"} 0.00
cpu_time_steal{cpu="cpu3"} 0.00
cpu_time_steal{cpu="cpu4"} 0.00
cpu_time_steal{cpu="cpu5"} 0.00
cpu_time_steal{cpu="cpu6"} 0.00
cpu_time_steal{cpu="cpu7"} 0.00
cpu_time_steal{cpu="cpu8"} 0.00
cpu_time_steal{cpu="cpu9"} 0.00
cpu_time_steal{cpu="cpu10"} 0.00
cpu_time_steal{cpu="cpu11"} 0.00
cpu_time_steal{cpu="cpu-total"} 0.00
# HELP cpu_time_guest Statistic CPU Guest time
# TYPE cpu_time_guest counter
cpu_time_guest{cpu="cpu0"} 0.00
cpu_time_guest{cpu="cpu1"} 0.00
cpu_time_guest{cpu="cpu2"} 0.00
cpu_time_guest{cpu="cpu3"} 0.00
cpu_time_guest{cpu="cpu4"} 0.00
cpu_time_guest{cpu="cpu5"} 0.00
cpu_time_guest{cpu="cpu6"} 0.00
cpu_time_guest{cpu="cpu7"} 0.00
cpu_time_guest{cpu="cpu8"} 0.00
cpu_time_guest{cpu="cpu9"} 0.00
cpu_time_guest{cpu="cpu10"} 0.00
cpu_time_guest{cpu="cpu11"} 0.00
cpu_time_guest{cpu="cpu-total"} 0.00
# HELP cpu_time_guest_nice Statistic CPU GuestNice time
# TYPE cpu_time_guest_nice counter
cpu_time_guest_nice{cpu="cpu0"} 0.00
cpu_time_guest_nice{cpu="cpu1"} 0.00
cpu_time_guest_nice{cpu="cpu2"} 0.00
cpu_time_guest_nice{cpu="cpu3"} 0.00
cpu_time_guest_nice{cpu="cpu4"} 0.00
cpu_time_guest_nice{cpu="cpu5"} 0.00
cpu_time_guest_nice{cpu="cpu6"} 0.00
cpu_time_guest_nice{cpu="cpu7"} 0.00
cpu_time_guest_nice{cpu="cpu8"} 0.00
cpu_time_guest_nice{cpu="cpu9"} 0.00
cpu_time_guest_nice{cpu="cpu10"} 0.00
cpu_time_guest_nice{cpu="cpu11"} 0.00
cpu_time_guest_nice{cpu="cpu-total"} 0.00
```

</details>

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add jadiunr/check-cpu-time
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/jadiunr/check-cpu-time].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-cpu-time
  namespace: default
spec:
  command: check-cpu-time --example example_arg
  subscriptions:
  - system
  runtime_assets:
  - jadiunr/check-cpu-time
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the check-cpu-time repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/jadiunr/check-cpu-time/blob/master/.github/workflows/release.yml
[5]: https://github.com/jadiunr/check-cpu-time/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
