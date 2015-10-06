# Update #2
In the past week I have made some progress with my CNF converter. The CNF converter is the piece of the project that will take the user's input files and convert them to my internal format for a Clause Set. The idea is to take a Clause given in plaintext such as `AvB` and to return a CNF form, `{A, B}`. When this conversion is done for each premise and the negation of the conclusion the set of all of the clauses returned clauses, the Clause Set, is passed to the Theorem Prover part of the project, which can then determine satisfiability.

So far I have the conversion from plaintext to my internal Clause Set format working for 'simple' statements that just include Ands, Ors, Literals, and Negations of Literals. This is a very simple subset of the actual work that will need to be done to have a full CNF converter, but it is a good proof of concept.

In the next week I hope to make more progress on the CNF converter by working on adding the ability to take implication and equivalence operators, substitute equivalent and/or operators and then convert that piece to CNF.


In addition, below is an example of an input file:
```
A
~A^B
B
```
In the input file format that the project uses, each line is a premise, except for the last, which is the conclusion.

```
A, B, ... == Literals
~         == Negation
^         == And
v         == Or
->        == Implication
<->       == Equivalence
```
