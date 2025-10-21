Given a string `str` and an array of positive integers `widths`, write a function that splits the string into lines, each with the exact number of characters as specified by the corresponding width.
Return an array of the substrings.
Use the last width for any remaining characters if the array is shorter than needed.

Example:

```js
const str = "Supercalifragilisticexpialidocious";
const widths = [5, 9, 4];

> splitByWidths(str, widths);
> ['Super', 'califragi', 'list', 'icex', 'pial', 'idoc', 'ious']
```
