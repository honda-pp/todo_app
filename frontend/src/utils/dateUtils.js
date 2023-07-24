export function convertSqlNullTimeToDateString(sqlNullTime) {
  // golang's sql.NullTime
  if (sqlNullTime.Valid) {
    const dateObj = new Date(sqlNullTime.Time);
    const year = dateObj.getFullYear();
    const month = String(dateObj.getMonth() + 1).padStart(2, '0');
    const day = String(dateObj.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }
  return null;
}
