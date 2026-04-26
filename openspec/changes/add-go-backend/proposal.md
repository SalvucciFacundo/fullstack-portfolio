# Propuesta: Backend en Go con Arquitectura Hexagonal para el Portfolio

## Intención y Alcance
El objetivo es profesionalizar el portfolio eliminando la data hardcodeada y reemplazándola por un backend dinámico y escalable. Vamos a construir una infraestructura que permita gestionar proyectos y experiencias de forma segura y eficiente.

El alcance del cambio incluye:
- Implementación de una API REST en Go.
- Base de datos Postgres para la persistencia.
- Autenticación administrativa mediante JWT.
- Inyección de dependencias con `samber/do`.
- Dockerización para despliegue en Oracle Cloud.

## Enfoque Técnico
- **Arquitectura Hexagonal (Ports & Adapters)**: Garantizamos que la lógica de negocio sea independiente de la base de datos y del framework de API.
- **Lenguaje**: Go 1.21+ para máxima performance y eficiencia.
- **DI**: `samber/do` para un manejo de dependencias limpio y moderno.
- **DB**: Postgres, aprovechando tipos de datos avanzados para optimizar las consultas de proyectos.
- **Auth**: Middleware de JWT para proteger las rutas administrativas.

## Impacto Esperado
- **Mantenibilidad**: Estructura clara que facilita la extensión y el testing.
- **Performance**: Respuesta ultra rápida del backend, ideal para una buena UX.
- **Seguridad**: Control total sobre quién puede modificar la información del portfolio.
- **Escalabilidad**: Preparado para crecer o cambiar de infraestructura sin dolor.
