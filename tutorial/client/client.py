import requests
from client_utils import build_query_string

BASE_URI = "/lala"


class Client:
    def __init__(self):
        self.url = BASE_URI

    def jobs_post(self, data, headers=None, query_params=None):
        """
        Create a job
        It is method for POST /jobs
        """
        uri = self.url + "/jobs"
        uri = uri + build_query_string(query_params)
        return requests.post(uri, data, headers=headers)

    def projects_post(self, data, headers=None, query_params=None):
        """
        Create a project
        It is method for POST /projects
        """
        uri = self.url + "/projects"
        uri = uri + build_query_string(query_params)
        return requests.post(uri, data, headers=headers)

    def resources_get(self, headers=None, query_params=None):
        """
        Get a resource
        It is method for GET /resources
        """
        uri = self.url + "/resources"
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)

    def resources_post(self, data, headers=None, query_params=None):
        """
        Post a resource
        It is method for POST /resources
        """
        uri = self.url + "/resources"
        uri = uri + build_query_string(query_params)
        return requests.post(uri, data, headers=headers)

    def resources_byResourceId_get(self, resourceId, headers=None, query_params=None):
        """
        Get a resource ID
        It is method for GET /resources/{resourceId}
        """
        uri = self.url + "/resources/"+resourceId
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)

    def resources_byResourceId_put(self, data, resourceId, headers=None, query_params=None):
        """
        Put resource ID
        It is method for PUT /resources/{resourceId}
        """
        uri = self.url + "/resources/"+resourceId
        uri = uri + build_query_string(query_params)
        return requests.put(uri, data, headers=headers)

    def resources_byResourceId_delete(self, resourceId, headers=None, query_params=None):
        """
        Delete a resource ID
        It is method for DELETE /resources/{resourceId}
        """
        uri = self.url + "/resources/"+resourceId
        uri = uri + build_query_string(query_params)
        return requests.delete(uri, headers=headers)

    def resources_byResourceId_yet_another_get(self, resourceId, headers=None, query_params=None):
        """
        Yet another
        It is method for GET /resources/{resourceId}/yet_another
        """
        uri = self.url + "/resources/"+resourceId+"/yet_another"
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)
