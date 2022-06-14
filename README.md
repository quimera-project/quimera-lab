<p align="center">
    <a href="https://github.com/quimera-project/quimera-lab/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/PwnedShell/Larascript?style=for-the-badge">
    </a>
    <a href="https://github.com/quimera-project/quimera-lab/network/members">
        <img alt="GitHub forks" src="https://img.shields.io/github/forks/quimera-project/quimera-lab?style=for-the-badge">
    </a>
    <a href="https://github.com/quimera-project/quimera-lab/stargazers">
        <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/quimera-project/quimera-lab?style=for-the-badge">
    </a>
    <a href="https://github.com/quimera-project/quimera-lab/issues">
        <img alt="GitHub issues" src="https://img.shields.io/github/issues/quimera-project/quimera-lab?style=for-the-badge">
    </a>
    <a href="https://github.com/quimera-project/quimera-lab/blob/main/LICENSE.md">
        <img alt="GitHub" src="https://img.shields.io/github/license/quimera-project/quimera-lab?style=for-the-badge">
    </a>
</p>

<p align="center">
  <a href="https://github.com/quimera-project/quimera-lab">
    <img src=".github/img/Quimera.jpg" alt="Quimera" width="25%"/>
  </a>

  <h1 align="center">Quimera Laboratory</h1>

  <p align="center">
    Quimera Laboratory proporciona un espacio de desarrollo para los binarios de los distintos checks, focalizado en Golang. Algunos de los checks necesitan apoyarse en un lenguaje como Go para conseguir un resultado mucho más eficaz y eficiente, con un código mucho más simple y organizado del que podría ofrecer un shell script. De todas formas, el repositorio puede almacenar, a su vez, binarios o scripts de otros lenguajes de programación, convirtiéndose así en un repositorio auxiliar para trabajar en la creación y testeo de binarios o scripts para los checks.
  </p>
</p>

> 🚧 Este proyecto proviene de un Trabajo Fin de Grado de Ingeniería de la Ciberseguridad. Se encuentra en una fase muy temprana del desarrollo y sufrirá diferentes cambios hasta llegar a una versión estable final.

## Instalación
```bash
git clone https://github.com/quimera-project/quimera-lab
```

## Uso

> Una vez creados y probados los ficheros correspondientes, estos deberán ser almacenados en Quimera Workshop. Quimera Laboratory no participa en el flujo de ejecución de Quimera y, por lo tanto, no es fundamental para el correcto funcionamiento de Quimera, siendo innecesario para el usuario final.

## Distribución

<p align="center">
    <img src=".github/img/Laboratory.jpg" alt="Laboratory" width="50%"/>
</p>


<p align="justify">
El repositorio se divide en dos carpetas: utils y checks. La primera presta distintas utilidades, relacionadas con los checks, a modo de paquete de Golang para los binarios creados en la carpeta checks. De esta manera, las funciones y variables más relevantes pueden ser reutilizadas por todos los binarios, sin repetir código. Por otro lado, la carpeta checks se subdivide en carpetas relacionadas con la categoría del check a la que pertenence el binario o script almacenado en ellas.
</p>
