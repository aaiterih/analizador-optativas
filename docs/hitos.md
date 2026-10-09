## Milestone 0: Modelo inicial del dominio
En este hito se modelan, aplicando diseño guiado por dominio (DDD), los conceptos que aparecen en la [HU001].

El trabajo se hace con issues que derivan directamente de los problemas definidos de la HU001, empezando por los objetos valor, y llevando el modelo hacia las partes en las que es obvio si son entidades u objetos valor. Paso a paso, en los propios issues, se va viendo si el modelo se está haciendo bien, y los issues se van ajustando con la retroalimentación continua. El código responde a ellos con commits que referencian al issue al que responden.

## Milestone 1: Paquete con tests automáticos
En este hito se desarrolla la lógica principal del proyecto a partir del paso anterior, y se incorpora la infraestructura de tests automáticos.

Para comprobar que este producto es válido, se usan las pruebas automáticas. El hito se considera cerrado cuando los tests (que validan la lógica de negocio y comprueban que efectivamente se resuelve el problema de una HU específica) se lanzan con la herramienta de tareas del proyecto y pasan en verde, y cada test se relaciona con el issue que comprueba.