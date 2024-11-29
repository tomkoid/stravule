export function formatDate(dateString: string) {
  let d = parseDate(dateString);
  let weekday = getWeekday(d);

  return `${weekday} ${d.getDate()}. ${d.getMonth() + 1}. ${d.getFullYear()}`;
}

export function parseDate(dateString: string): Date {
  // const czechDateFormat = new Intl.DateTimeFormat("cs-CZ");
  // let date = new Date(czechDateFormat.formatToParts(new Date(dateString)).map(part => part.value).join('/'));
  const [day, month, year] = dateString.split(".").map(part => parseInt(part, 10));

  let date = new Date(year, month - 1, day);
  if (isNaN(date.getTime())) {
    console.error("parsed date is undefined, date string:", dateString);
    date = new Date(dateString + " GMT+1");
  }
  return date;
}

export function getWeekday(date: Date) {
  const days = ["Neděle", "Ponděli", "Úterý", "Středa", "Čtvrtek", "Pátek", "Sobota"];
  return days[date.getDay()];
}
