**How to install/create a package:**

 1. Create a file inside `src/`
 2. Define it as a package in the top line of the file `package pkgName;`
 3. Run `go install pkgName` (do this for every change to the file)
 4. You can view the compiled version in `pkg/"system_arch"/`

*Note however for a function to be used inside an import it must start with a capital letter. (This defines that it is exportable.)*

**Importing the package:**

 Either add to an existing import chain like so:

    import (
		"fmt";
		div "divider";
	);

Or, directly import it like so: `import pkgShorthand "pkgName";`