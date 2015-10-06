# Update #1 - Go for Gold
I am very excited to be starting my second semester working on a Davis Putnam Automated Theorem Prover Project. Last semester, I implemented the algorithm as a final project for the class Computability and Logic here at RPI. The code for this project can be found [here!](https://github.com/astonshane/davisputnam). I learned a lot about automated theorem proving in general and the Davis-Putnam algorithm while working on that project and I am hoping to expand upon this knowledge this semester in RCOS.

My first major goal for the project this semester has been to move the project from Python to Go. I choose to move the project to Go for a number of reasons. First off, in the later stages of the Python based project from last semester, I started to hit some road bumps in terms of performance. The Davis-Putnam Algorithm works great on small Clause-Set sizes, but as you start to deal with more and more clauses and more and more literals inside those clauses, the algorithm quickly becomes very computationally intensive. My Python implementation was taking multiple seconds per run on moderately sized test cases. I decided that I should move to a language with a heavier focus on performance.

There are a number of language benchmarking tools available online such as [this](http://benchmarksgame.alioth.debian.org/u32/performance.php?test=nbody) one. After consulting many I settled upon Go. But "Why?", you may ask? If ALL I was after was extreme performance gains I would have ended up with C++ or \**shuders*\* Fortran. Go offers significant speed improvements over Python but remains much more expressive than C++ or Fortran. In this way, it is a nice middle ground.

I am happy to say that so far in the semester I have made a lot of progress in converting the project to Go. I have already completed the porting of the Satisfiable() function that makes up the "meat" of the Algorithm. Next up, will be the creation of a Conjunctive Normal Form converter. Look for a post about that soon.