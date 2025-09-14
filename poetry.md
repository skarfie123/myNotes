# Poetry

- generate `pyproject.toml`: `poetry init`
  - with defaults: `poetry init -n`
- generate new project: `poetry new <path>`
- display config `poetry config --list`
- set venvs to be in project: `poetry config virtualenvs.in-project true`
  - this is helps vscode detect the venv
- add dependency: `poetry add <dep>`
- add dev dependency: `poetry add pytest --group dev`
  - also make the group optional

    ```toml
    [tool.poetry.group.dev]
    optional = true
    ```

- list dependencies: `poetry show`
  - in tree view: `poetry show --tree`
- spawn virtual environment shell: `poetry shell`
  - need to include plugin: <https://github.com/python-poetry/poetry/≥pull/9763>
- run a command in the environment: `poetry run <command>`
- install dependencies: `poetry install`
- update dependencies: `poetry update`
- export the lockfile for CI: `poetry export -f requirements.txt --output requirements.txt --without-hashes`
  - need to include plugin: <https://github.com/python-poetry/poetry/pull/5980>
- configure PyPI: `poetry config pypi-token.pypi <my-token>`
  - <https://python-poetry.org/docs/repositories/#configuring-credentials>
- publish: `poetry publish --build`
