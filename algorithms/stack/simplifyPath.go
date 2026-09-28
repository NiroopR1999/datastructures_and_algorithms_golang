package stack

import "strings"

func simplifyPath(path string) string {

	// WHY:
	// Split the path into individual components using "/".
	//
	// Example:
	// "/home//foo/../bar"
	//
	// becomes:
	// ["", "home", "", "foo", "..", "bar"]
	parts := strings.Split(path, "/")

	// WHY:
	// We use a stack because:
	// - A normal directory means "go into this directory" → PUSH
	// - ".." means "go back to the parent directory" → POP
	stack := []string{}

	for _, part := range parts {

		// WHY:
		// "" is created by consecutive "/" characters.
		// "." means the current directory.
		//
		// Neither changes our current location, so ignore them.
		if part == "" || part == "." {
			continue
		}

		// WHY:
		// ".." means "go to the parent directory".
		// Therefore, remove the most recently added directory.
		if part == ".." {

			// WHY:
			// If the stack is empty, we are already at the root "/".
			// There is no directory above root, so ".." does nothing.
			if len(stack) != 0 {

				// WHY:
				// Remove the last element from the stack.
				//
				// Example:
				// [home user docs]
				//        ↓ ".."
				// [home user]
				stack = stack[:len(stack)-1]
			}

			// WHY:
			// ".." has already been handled.
			// Without continue, it could fall through and get
			// incorrectly pushed onto the stack.
			continue
		}

		// WHY:
		// Anything else is a valid directory/file name,
		// so push it onto the stack.
		//
		// IMPORTANT:
		// "..." and "...." are normal names.
		// Only "." and ".." have special meanings.
		stack = append(stack, part)
	}

	// WHY:
	// The stack now contains the directories that should remain
	// in the canonical path.
	//
	// Example:
	// [home user docs]
	//
	// strings.Join → "home/user/docs"
	//
	// Add "/" because the path must always be absolute.
	return "/" + strings.Join(stack, "/")
}

/*
===========================================================
ENTIRE LOGIC TO REMEMBER
===========================================================

Split path
    ↓
Process each part
    ↓
"" or "."  → IGNORE
    ↓
".."       → POP
    ↓
anything else → PUSH
    ↓
Join stack with "/"
    ↓
Add "/" at beginning


ONE RULE THAT PREVENTS CONFUSION WITH DOTS:

"."    → special → IGNORE
".."   → special → POP
"..."  → normal name → PUSH
"...." → normal name → PUSH
===========================================================
*/