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

### Usage of `env.gocyclo_report`

The `env.gocyclo_report` is an environment variable that holds the results of the `gocyclo` analysis. This report includes details about the cyclomatic complexity of each function in your Go code.

Here's how you can use it:

1. **Accessing the report**: You can access the `gocyclo_report` in your workflow file using the `env` context. For example, you can print the report to the console with the following step:

```github-actions-workflow
- name: Print gocyclo report
  run: echo "${{ env.gocyclo_report }}"
```

2. **Using the report in other steps:** The gocyclo_report can be used in other steps of your workflow. For example, you can use it to conditionally run steps based on the complexity of your code, or to post the report as a comment in a pull request.

3. **Storing the report:** If you want to store the gocyclo_report for future reference, you can write it to a file in your repository. For example:

```github-actions-workflow
- name: Store gocyclo report
  run: echo "${{ env.gocyclo_report }}" > gocyclo_report.txt
```

Remember, the gocyclo_report is a powerful tool for understanding the complexity of your Go code. Use it wisely to maintain the quality of your codebase.

## Usage

Below is an example of how to use this action in your workflow:

```yaml
name: Analyze Go Code with gocyclo

on:
  pull_request:
    paths:
      - "**/*.go"

# Permissions for marocchino/sticky-pull-request-comment@v2
permissions:
  pull-requests: write

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
            ${{ env.gocyclo_report }}
```

### About gocyclo

The gocyclo tool, created by [fzipp](https://github.com/fzipp), is a command-line utility that computes the cyclomatic complexity of functions in Go source code. Cyclomatic complexity is a metric that measures the number of linearly independent paths through a program's source code, which is a useful indicator of code complexity and maintainability.

For more information about `gocyclo` , visit the [gocyclo GitHub repository](https://github.com/fzipp/gocyclo).

By integrating gocyclo into your GitHub Actions workflow, you can continuously monitor and manage the complexity of your Go codebase, ensuring that your software remains maintainable and scalable.
