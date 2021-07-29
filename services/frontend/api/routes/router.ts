import { Router } from 'express'
import instance from '../instance'

const router = Router()

router.get('/capes', (req, res) => {
  (async () => {
    try {
      let url: string = "capes"
      if (req.query.form === 'true') {
        url = "capes?form=true"
      }
      const api = await instance.get(url);
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }

  })();
});

router.put('/capes', (req, res) => {
  (async () => {
    try {
      const api = await instance.put('capes/', req.body, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.delete('/capes/:cape', (req, res) => {
  (async () => {
    try {
      const api = await instance.delete('capes/' + req.params.cape, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/capes/:cape', (req, res) => {
  (async () => {
    try {
      const api = await instance.put('capes/' + req.params.cape, req.body, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.get('/projects', (req, res) => {
  (async () => {
    try {
      const api = await instance.get("projects");
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.get('/projects/:project', (req, res) => {
  (async () => {
    try {
      const api = await instance.get("projects/" + req.params.project);
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/projects/:project', (req, res) => {
  (async () => {
    try {
      const api = await instance.put("projects/" + req.params.project, req.body, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/projects/:project/link', (req, res) => {
  (async () => {
    try {
      const api = await instance.put('projects/' + req.params.project + '/link', req.body, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.get("/auth/users", (req, res) => {
  (async () => {
    try {
      const api = await instance.get("auth/users");
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/auth/users/:uuid', (req, res) => {
  (async () => {
    try {
      const api = await instance.put('auth/users/' + req.params.uuid, req.body);
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/auth/login/', (req, res) => {
  (async () => {
    try {
      const api = await instance.put('auth/login/', req.body);
      res.status(api.status).send(api.data);
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/auth/register/', (req, res) => {
  (async () => {
    try {
      const api = await instance.put('auth/register/', req.body);
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }

  })();
});

router.get('/auth/token/', (req, res) => {
  (async () => {
    try {
      const api = await instance.get('auth/token/', {headers: {authorization: req.headers.authorization}});
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }

  })();
});

router.get('/auth/me/', (req, res) => {
  (async () => {
    try {
      const api = await instance.get('auth/me/', { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.post('/auth/users/:uuid/admin', (req, res) => {
  (async () => {
    try {
      const api = await instance.post('auth/users/' + req.params.uuid + '/admin', req.body, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.post('/auth/users/:uuid/hwid', (req, res) => {
  (async () => {
    try {
      const api = await instance.post('auth/users/' + req.params.uuid + '/hwid', req.body, {headers: {authorization: req.headers.authorization}});
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }

  })();
});

router.post('/auth/users/:uuid/access', (req, res) => {
  (async () => {
    try {
      const api = await instance.post('auth/users/' + req.params.uuid + '/access', req.body, {headers: {authorization: req.headers.authorization}});
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }

  })();
});

router.delete('/me/online', (req, res) => {
  (async () => {
    try {
      const api = await instance.delete("me/online", { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.put('/me/online', (req, res) => {
  (async () => {
    try {
      const api = await instance.put("me/online", req.body, { headers: { authorization: req.headers.authorization } });
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

router.get('/me/online', (req, res) => {
  (async () => {
    try {
      const api = await instance.get("me/online");
      res.status(api.status).send(api.data)
    } catch (e) {
      res.status(e.response.status).send(e.response.data);
    }
  })();
});

export default router
