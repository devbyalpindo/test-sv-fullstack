import callAPI from "../helpers/callApi";

const baseUrl = "http://test-sv:8080";

export async function getAllArticle() {
  const url = `${baseUrl}/article`;

  return callAPI({
    url,
    method: "GET",
  });
}

export async function getAllArticleWithLimit(limit: number, offset: number) {
  const url = `${baseUrl}/article/${limit}/${offset}`;

  return callAPI({
    url,
    method: "GET",
  });
}

export async function addArticle(payload: string) {
  const url = `${baseUrl}/article`;

  return callAPI({
    url,
    method: "POST",
    data: payload,
  });
}

export async function editArticle(id: string, payload: string) {
  const url = `${baseUrl}/article/${id}`;

  return callAPI({
    url,
    method: "PUT",
    data: payload,
  });
}

export async function getArticleById(id: string) {
  const url = `${baseUrl}/article/${id}`;

  return callAPI({
    url,
    method: "GET",
  });
}

export async function removeArticleById(id: number) {
  const url = `${baseUrl}/article/${id}`;

  return callAPI({
    url,
    method: "DELETE",
  });
}
