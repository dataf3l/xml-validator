# xml-validator

The purpose of this program is to verify from a certain path, all the files that
are in that folder, to determine if these files are a valid .gz file. To determine
this, the first 2 bytes of each file, if equal to [31  139 ... ], then it is a valid .gz file.

Returns a string with all paths to valid files repaired by commas.

### References
