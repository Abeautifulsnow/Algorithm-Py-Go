const source = [
  { id: 19, parentId: 0 },
  { id: 18, parentId: 16 },
  { id: 17, parentId: 16 },
  { id: 16, parentId: 0 },
  { id: 15, parentId: 17 },
  { id: 14, parentId: 19 },
  { id: 21, parentId: 19 },
  { id: 20, parentId: 18 },
];

function convertM1(source, parentKey, currentKey, rootValue) {
  const hash_res = {};

  for (let item of source) {
    const parentV = item[parentKey];
    const children = hash_res[parentV] || [];

    if (children.length === 0) {
      children.push(item);
      hash_res[parentV] = children;
    } else {
      children.push(item);
    }
  }

  const recursion = (children) => {
    const tmp = [];

    for (let cItem of children) {
      const v = hash_res[cItem[currentKey]];

      if (!v) {
        tmp.push(cItem);
      } else {
        cItem["children"] = recursion(v);
        tmp.push(cItem);
      }
    }

    return tmp;
  };

  const outResult = {};
  const targetValue = hash_res[rootValue] || [];

  if (targetValue) {
    outResult[currentKey] = rootValue;
    outResult["children"] = recursion(targetValue);
  }

  console.log(JSON.stringify(outResult, "\t", 2));
}

function convertM2(source, parentKey, currentKey, rootValue) {
  const outRes = {};

  for (let i = 0; i < source.length; i++) {
    const item = source[i];
    if (Number(item[parentKey]) === Number(rootValue)) {
      outRes[currentKey] = rootValue;
      outRes["children"] = outRes["children"] || [];
      outRes["children"].push(item);
      source.splice(i, 1);
      i--;
    }
  }

  const restSource = source;
  const recursion = (children) => {
    if (children) {
      for (let cItem of children) {
        for (let j = 0; j < restSource.length; j++) {
          const jItem = restSource[j];

          if (Number(cItem[currentKey]) === Number(jItem[parentKey])) {
            const subChildren = cItem["children"] || [];
            subChildren.push(jItem);
            cItem["children"] = subChildren;
            restSource.splice(j, 1);
            j--;
          } else {
            recursion(cItem["children"]);
          }
        }
      }
    }
  };

  recursion(outRes["children"]);

  console.log(JSON.stringify(outRes, "\t", 2));
}

function convertM3(source, parentKey, currentKey, rootValue) {
  const res = source.reduce(
    (p, c) => {
      c = Object.assign((p[c[currentKey]] ??= {}), c);
      ((p[c[parentKey]] ??= {}).children ??= []).push(c);
      return p;
    },
    { rootValue: { currentKey: rootValue } }
  );

  return res[0];
}

function convertM4(source, parentKey, currentKey, rootValue) {
  const dfs = (rootV) => {
    const childrenRes = [];

    source.forEach((item) => {
      if (item[parentKey] === rootV) {
        const subChildren = dfs(item[currentKey]);
        subChildren.length > 0
          ? childrenRes.push({ ...item, children: subChildren })
          : childrenRes.push(item);
      }
    });

    return childrenRes;
  };

  return {
    [currentKey]: rootValue,
    children: dfs(rootValue),
  };
}

console.log(JSON.stringify(convertM4(source, "parentId", "id", 0), "\t", 2));

// output

`
  {
    "id": 0,
    "children": [
      {
        "id": 19,
        "parentId": 0,
        "children": [
          {
            "id": 14,
            "parentId": 19
          },
          {
            "id": 21,
            "parentId": 19
          }
        ]
      },
      {
        "id": 16,
        "parentId": 0,
        "children": [
          {
            "id": 18,
            "parentId": 16,
            "children": [
              {
                "id": 20,
                "parentId": 18
              }
            ]
          },
          {
            "id": 17,
            "parentId": 16,
            "children": [
              {
                "id": 15,
                "parentId": 17
              }
            ]
          }
        ]
      }
    ]
  `;
