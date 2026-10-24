export const formatDate = (date: Date) => {
  const day = date.getUTCDate()
  const month = date.toLocaleDateString('en-US', {
    month: 'short',
    timeZone: 'UTC',
  })
  const year = date.getUTCFullYear()
  return `${day} ${month} ${year}`
}
