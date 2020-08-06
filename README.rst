=============================================
``poplar`` -- Golang Performance Test Toolkit
=============================================

Overview
--------

Poplar is a set of tools for running and recording results for
performance tests suites and benchmarks. It provides easy integration
with a number of loosely related tools:

- `ftdc <https://github.com/deciduosity/ftdc>`_ is a compression format for
  structured and semi-structured timeseries data. Poplar provides
  service integration for generating these data payloads.

- `cedar <https://github.com/evergreen-ci/cedar>`_ is a service for collecting
  and processing data from builds. Poplar provides a client for uploading test
  results to cedar from static YAML or JSON data. There are no compile-time
  dependencies on Cedar.

Additionally, poplar provides a complete benchmark test harness with
integration for collecting ftdc data and sending that data to cedar,
or reporting it externally. 

Poplar is useful for building and capturing benchmarks independently of the
inter-run analysis provided by cedar.

Some popular functionality is included as a CLI in the `curator
<https://github.com/mongodb/curator>`_ tool, as ``curator poplar``.


Documentation
-------------

See the `godoc <https://godoc.org/github.com/decidousity/poplar/>`_
for complete documentation. 

Use
---

If you encounter a problem or have a feature that you'd like to see added to
``mrpc``, please feel free to create an issue or file a pull request.

poplar is available for use under the terms of the Apache License (v2). 

Development
-----------

This version of popular (deciduosity) provides
compatibility with more recent versions of Go and its module system. 

All project automation is managed by a makefile, with all output captured in the
`build` directory. Consider the following operations: ::

   make build                   # runs a test compile
   make test                    # tests all packages
   make lint                    # lints all packages
   make test-<package>          # runs the tests only for a specific packages
   make lint-<package>          # lints a specific package
   make html-coverage-<package> # generates the coverage report for a specific package
   make coverage-html           # generates the coverage report for all packages

The buildsystem also has a number of flags, which may be useful for more
iterative development workflows: ::

  RUN_TEST=<TestName>   # specify a test name or regex to run a subset of tests
  RUN_COUNT=<num>       # run a test more than once to isolate an intermittent failure
  RACE_DETECTOR=true    # run specified tests with the race detector enabled. 
