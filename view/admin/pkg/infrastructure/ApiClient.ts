export class ApiClient {
  async get(url: string, token: string): Promise<any> {
    return await $fetch(url, {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    });
  }

  async post(url: string, data: any, token: string): Promise<any> {
    return await $fetch(url, {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: data,
    });
  }

  async put(url: string, data: any, token: string): Promise<any> {
    return await $fetch(url, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: data,
    });
  }

  async delete(url: string, token: string): Promise<any> {
    return await $fetch(url, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + token,
      },
    });
  }
}

const apiClient = new ApiClient();
export default apiClient;
