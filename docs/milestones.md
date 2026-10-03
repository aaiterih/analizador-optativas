## Milestone 0: Modelo inicial del dominio
En este hito se entrega el código que modela los conceptos que aparecen en la [HU001].

Para comprobar que esta primera versión es válida, se evalúa el proceso seguido, que es la aplicación de diseño guiado por dominio (DDD). El hito se considera cerrado cuando los issues derivan directamente de los problemas definidos de la HU001, en cada issue se decide si un concepto es entidad u objeto valor, y el código responde a ellos con commits que referencian al issue al que responden.

## Milestone 1: Paquete con tests automáticos
En este hito se desarrolla la lógica principal del proyecto a partir del paso anterior, y se incorpora la infraestructura de tests automáticos.

Para comprobar que este producto es válido, se usan las pruebas automáticas. El hito se considera cerrado cuando los tests (que validan el correcto funcionamiento de las entidades y objetos valor) se lanzan con la herramienta de tareas del proyecto y pasan en verde, y cada test se relaciona con el issue que comprueba