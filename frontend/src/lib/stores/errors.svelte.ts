export function createErrors() {
  let errors: string[] = $state([]);

  return {
    get errors() { return errors },
    get length() { return errors.length },
    clean: () => { errors.shift() },
    removeOne: () => { errors.pop() },
    add: (error: string) => { errors.push(error) },
  };
}

export let errors = createErrors()
