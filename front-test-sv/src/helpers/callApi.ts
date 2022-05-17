interface CallApiProps {
  url: string;
  method: string;
  data?: string;
}

export default async function callAPI({
  url,
  method,
  data,
}: CallApiProps) {
  let headers: any = {
    "Content-Type": "application/json",
    "cache-control": "no-cache",
  };

  const response = await fetch(url, {
    method,
    headers,
    body: data,
  });

  const result = await response.json();
  return result;
}
