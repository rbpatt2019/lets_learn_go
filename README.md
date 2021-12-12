# lets_learn_go

Because we should always be learning something new!

## Motivation

Simply,
to expand my coding knowledge!
Coming from a background of interpreted languages (Python, R),
compiled languages represent a bit of a gap in my knowledge.
Though there are many to choose from,
[Go][golang] struck me as a good choice. 
It is open-sourced,
actively developed,
natively support concurrency,
has garbage collections,
and is rapidly becoming a dominant language for cloud services.
One need only look at the list of companies using it!

## Plan

The code here in will be guided by John Bodner's excellent book,
[Learning Go][learning_go].
Please support him and pick up a copy at the above link!
**(NOT SPONSORED)**

## Other Notes

I'm also going to be trying out a few new tools.
Instead of the usual `Makefile`,
I'm going to try a [Justfile][just],
It seems a much more intuitive,
polished version of the venerable `Make`.
I've also tried moving my hooks into a `.githooks` directory.
This way,
the necessary hooks and message template can be provided without needing external tools.
Since the setup can be handled by `just`,
this steamlines development for any future collaborators.
Win-win!

[golang]: https://go.dev "GoLang"
[learning_go]: https://www.oreilly.com/library/view/learning-go/9781492077206/ "Learning Go"
[just]: https://github.com/casey/just "Just"
