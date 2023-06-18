## Fiat Shamir Non-Interactive Zero Knowledge Proof with an Elliptic Curve Implementation

Curve used here is Edwards25519

To run the project

Run this command:

```
$ go build
```

For Docker image

Run this command

```
$ docker compose up -d
```

Sample Run

Bob and Alice agree:
G: 3fd96d7b139e8805e3d94f1e04cd637aae7b9c2565a708464b62449cbbfb07fa
H: f7bf1938a30b2a84e745cbe96b25854be4a62ea69b936b8071bad412c89118f5

Bob's Password: qwerty
Bob's Secret (x): 49f9c587f98c1e5840ee76f205437cf9a582b27168c0ea744b2cf58ee0233705

Bob sends these values:
xG: b21b0fe0af6c874300fd3dcd95b6c7a2d98ef9014632661a9f86e16e134cbc0d
xH 6d009dacea721114976c5da13fbe88afeaac8ff204dfdf73b6a5c4f9b1bb651e

Alice sends challenge:
c: 7592f5abb87a6797e6f63eb0a571a51a0197a258d25002079acd82e5dc9c990e

Bob computes:
v: 18025bbd0c975f8762a263603ab1c1887884e10a748a65b2f2a428a90994220f
r: af90b47cf39b03966bd5003de377147c0103e4d22b424740a23dcd046057300e

Proof correct

Explanation
