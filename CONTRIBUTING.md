# Contributing to Go Clean Architecture Template

We love your input! We want to make contributing to this template as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Code Style Guidelines

### Go Style
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt`
- Follow clean architecture principles
- Write meaningful commit messages
- Include comments for exported functions and packages

### Project Structure
```
.
├── internal/
│   ├── domain/         # Business logic and interfaces
│   ├── application/    # Use cases
│   ├── infrastructure/ # External implementations
│   └── interface/      # HTTP handlers
├── prisma/            # Database schema
└── scripts/           # Utility scripts
```

### Clean Architecture Rules
1. Inner layers don't know about outer layers
2. Domain entities contain no external dependencies
3. Use interfaces for dependency inversion
4. Infrastructure implements interfaces defined in domain

## Testing Standards

1. **Unit Tests**
   - Test business logic in isolation
   - Use mocks for external dependencies
   - Aim for high coverage in domain and application layers

2. **Integration Tests**
   - Test complete workflows
   - Use test database
   - Test API endpoints end-to-end

3. **Running Tests**
```bash
# Run all tests
make test

# Run with coverage
make test-coverage
```

## Pull Request Process

1. Update the README.md with details of changes if API is changed.
2. Update the CHANGELOG.md with your changes.
3. The PR will be merged once you have the sign-off of at least one maintainer.

## Any contributions you make will be under the MIT Software License
In short, when you submit code changes, your submissions are understood to be under the same [MIT License](LICENSE) that covers the project. Feel free to contact the maintainers if that's a concern.

## Report bugs using GitHub's [issue tracker]
We use GitHub issues to track public bugs. Report a bug by [opening a new issue]().

## Write bug reports with detail, background, and sample code

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can.
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

## License
By contributing, you agree that your contributions will be licensed under its MIT License.