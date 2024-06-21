# gocyclo Analysis Reports

Analyze Go files with gocyclo and report cyclomatic complexity.

## Overview

The `gocyclo Analysis Reports` GitHub Action allows you to analyze the cyclomatic complexity of your Go code using the `gocyclo` tool. This action helps developers maintain and improve the quality of their code by identifying functions with high cyclomatic complexity, which can be difficult to understand, test, and maintain.

## Purpose

This action is designed to:
- Analyze the cyclomatic complexity of Go functions.
- Provide detailed reports on the complexity of your code.
- Help identify high-risk areas that may require refactoring.

## Target Users

This action is ideal for:
- Go developers who want to maintain clean and maintainable code.
- Teams that follow continuous integration practices and want to ensure code quality.
- Developers who want to integrate cyclomatic complexity analysis into their CI/CD pipelines.

## Inputs

The action supports the following inputs:

- `ignore_pattern`: Pattern to ignore files in gocyclo analysis (default: '')
- `over`: Show functions with complexity > N only and return exit code 1 if the set is non-empty (default: '')
- `top`: Show the top N most complex functions only (default: '')

## Usage

Below is an example of how to use this action in your workflow:

```yaml
name: Analyze Go Code with gocyclo

on:
  pull_request:
    paths:
      - "**/*.go"

jobs:
  gocyclo_analysis:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4.0.0

      - name: Run gocyclo Analysis
        uses: snkrheadz/gocyclo-analysis-reports@v1
        id: gocyclo_analysis
        with:
          ignore_pattern: '_test|_mock|mock_|.pb.go|proto'
          over: '20'
          top: '10'
      - name: Post Results
        uses: marocchino/sticky-pull-request-comment@v2
        with:
          recreate: true
          header: gocyclo Report
          message: |
            ${{ steps.gocyclo_analysis.outputs.report }}
```
