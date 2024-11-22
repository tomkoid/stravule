let _rune = $state(false)

export const pageLoading = {
  get value() {
    return _rune
  },
  set value(newVal) {
    _rune = newVal
  }
}


let _pickOrdersRune: boolean = $state(false)

export const pickOrders = {
  get value() {
    return _pickOrdersRune
  },
  set value(newVal) {
    _pickOrdersRune = newVal
  }
}
