let _rune = $state(false)

export const pageLoading = {
  get value() {
    return _rune
  },
  set value(newVal) {
    _rune = newVal
  }
}
