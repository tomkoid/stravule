interface Error {
  value: string,
  createTime: Date
}

export function createErrors() {
  let errorsVar: Error[] = $state([]);

  return {
    get errors() {
      let errorList: string[] = []

      for (let i = 0; i < errorsVar.length; i++) {
        errorList.push(errorsVar[i].value)
      }

      return errorList
    },
    get length() { return errorsVar.length },
    clean: () => { errorsVar.shift() },
    cleanExpired: () => { errorsVar = errorsVar.filter(e => e.createTime.getTime() > (new Date().getTime() - 5000)) },
    removeOne: () => { errorsVar.pop() },
    add: (error: string) => {
      errorsVar.push({
        value: error,
        createTime: new Date()
      })

      const delay = 5000 //ms
      new Promise((r: any) => setTimeout(r, delay)).then(() => {
        errorsVar = errorsVar.filter(e => e.createTime.getTime() > (new Date().getTime() - delay))
      });
    },
  };
}

export let errors = createErrors()
