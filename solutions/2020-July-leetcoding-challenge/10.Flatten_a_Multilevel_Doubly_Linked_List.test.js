const target = require("./10.Flatten_a_Multilevel_Doubly_Linked_List");
const { arrToNodes, nodesToArr } = require("../../util/doubly-linked-list");

describe("2020-July-leetcoding-challenge 10", () => {
  test("case 1", () => {
    const head = arrToNodes([1, 2, null, 3]);
    const result = target(head);
    expect(nodesToArr(result)).toEqual([1, 3, 2]);
  });

  test("case 2", () => {
    const head = arrToNodes([
      1,
      2,
      3,
      4,
      5,
      6,
      null,
      null,
      null,
      7,
      8,
      9,
      10,
      null,
      null,
      11,
      12,
    ]);
    const result = target(head);
    expect(nodesToArr(result)).toEqual([1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6]);
  });

  test("case 3", () => {
    const head = arrToNodes([
      4201,
      null,
      1937,
      null,
      3203,
      null,
      4328,
      null,
      4557,
      null,
      880,
      null,
      273,
      null,
      3563,
      null,
      4440,
      null,
      1266,
      null,
      1890,
      null,
      4884,
      null,
      3128,
      null,
      2645,
      null,
      2681,
      null,
      3440,
      null,
      3826,
      null,
      2883,
      null,
      2924,
      null,
      726,
      null,
      2093,
      null,
      214,
      null,
      3203,
      null,
      3799,
      null,
      2145,
      null,
      1207,
      null,
      3714,
      null,
      4747,
      null,
      2618,
      null,
      2295,
      null,
      117,
      null,
      285,
      null,
      2296,
      null,
      4779,
      null,
      2789,
      null,
      42,
      null,
      2803,
      null,
      834,
      null,
      1464,
      null,
      1823,
      null,
      2190,
      null,
      1315,
      null,
      4398,
      null,
      624,
      null,
      588,
      null,
      4560,
      null,
      2890,
      null,
      4514,
      null,
      441,
      null,
      4251,
      null,
      3552,
      null,
      712,
      null,
      3381,
      null,
      541,
      null,
      3541,
      null,
      665,
      null,
      1646,
      null,
      950,
      null,
      1789,
      null,
      3835,
      null,
      722,
      null,
      4419,
      null,
      3761,
      null,
      3506,
      null,
      2823,
      null,
      3160,
      null,
      2399,
      null,
      2589,
      null,
      1189,
      null,
      155,
      null,
      3903,
      null,
      1438,
      null,
      4047,
      null,
      2566,
      null,
      558,
      null,
      4837,
      null,
      4313,
      null,
      1339,
      null,
      1211,
      null,
      2186,
      null,
      2404,
      null,
      1691,
      null,
      1651,
      null,
      407,
      null,
      1119,
      null,
      1476,
      null,
      158,
      null,
      3580,
      null,
      3824,
      null,
      4455,
      null,
      3595,
      null,
      1837,
      null,
      1602,
      null,
      2900,
      null,
      1907,
      null,
      408,
      null,
      4817,
      null,
      2485,
      null,
      2824,
      null,
      2936,
      null,
      192,
      null,
      1410,
      null,
      1613,
      null,
      491,
      null,
      4431,
      null,
      1522,
      null,
      1886,
      null,
      724,
      null,
      2046,
      null,
      3589,
      null,
      3658,
      null,
      682,
      null,
      1106,
      null,
      2215,
      null,
      4499,
      null,
      286,
      null,
      4852,
      null,
      1463,
      null,
      1964,
      null,
      4976,
      null,
      3936,
      null,
      2562,
      null,
      3839,
      null,
      1017,
      null,
      966,
      null,
      3021,
      null,
      2168,
      null,
      1195,
      null,
      3883,
      null,
      1171,
      null,
      4809,
      null,
      1438,
      null,
      1764,
      null,
      4181,
      null,
      4167,
      null,
      4388,
      null,
      2919,
      null,
      3745,
      null,
      4239,
      null,
      3181,
      null,
      3754,
      null,
      4228,
      null,
      1282,
      null,
      3458,
      null,
      4216,
      null,
      2034,
      null,
      2956,
      null,
      1539,
      null,
      728,
      null,
      4830,
      null,
      219,
      null,
      4544,
      null,
      110,
      null,
      4202,
      null,
      3335,
      null,
      4304,
      null,
      4146,
      null,
      3726,
      null,
      1692,
      null,
      1087,
      null,
      858,
      null,
      2911,
      null,
      3103,
      null,
      346,
      null,
      289,
      null,
      1209,
      null,
      1274,
      null,
      1759,
      null,
      4279,
      null,
      671,
      null,
      567,
      null,
      3136,
      null,
      4636,
      null,
      439,
      null,
      3359,
      null,
      1159,
      null,
      59,
      null,
      245,
      null,
      1290,
      null,
      3473,
      null,
      3796,
      null,
      2783,
      null,
      2536,
      null,
      1657,
      null,
      1747,
      null,
      1449,
      null,
      528,
      null,
      1318,
      null,
      4648,
      null,
      1477,
      null,
      2219,
      null,
      4050,
      null,
      566,
      null,
      1677,
      null,
      565,
      null,
      3807,
      null,
      561,
      null,
      4627,
      null,
      2176,
      null,
      67,
      null,
      4763,
      null,
      4796,
      null,
      2367,
      null,
      167,
      null,
      1767,
      null,
      4661,
      null,
      3322,
      null,
      4275,
      null,
      4424,
      null,
      356,
      null,
      4688,
      null,
      1480,
      null,
      4060,
      null,
      3910,
      null,
      3618,
      null,
      3751,
      null,
      350,
      null,
      4075,
      null,
      2493,
      null,
      4700,
      null,
      1808,
      null,
      1678,
      null,
      3927,
      null,
      779,
      null,
      4750,
      null,
      820,
      null,
      1422,
      null,
      3969,
      null,
      3017,
      null,
      1557,
      null,
      3886,
      null,
      2110,
      null,
      2735,
      null,
      4436,
      null,
      3170,
      null,
      3683,
      null,
      4447,
      null,
      2106,
      null,
      1785,
      null,
      707,
      null,
      354,
      null,
      1312,
      null,
      843,
      null,
      1558,
      null,
      17,
      null,
      4069,
      null,
      3281,
      null,
      1742,
      null,
      3048,
      null,
      1276,
      null,
      4357,
      null,
      870,
      null,
      3425,
      null,
      4149,
      null,
      668,
      null,
      4709,
      null,
      4415,
      null,
      442,
      null,
      1105,
      null,
      3801,
      null,
      4657,
      null,
      2940,
      null,
      273,
      null,
      2855,
      null,
      3971,
      null,
      2878,
      null,
      3569,
      null,
      218,
      null,
      4406,
      null,
      3372,
      null,
      471,
      null,
      2635,
      null,
      2010,
      null,
      1842,
      null,
      762,
      null,
      1477,
      null,
      2401,
      null,
      3152,
      null,
      1146,
      null,
      2492,
      null,
      3786,
      null,
      1073,
      null,
      2804,
      null,
      3354,
      null,
      877,
      null,
      2951,
      null,
      4426,
      null,
      3166,
      null,
      3728,
      null,
      2187,
      null,
      1859,
      null,
      3572,
      null,
      897,
      null,
      3276,
      null,
      3242,
      null,
      4024,
      null,
      4906,
      null,
      6,
      null,
      2708,
      null,
      854,
      null,
      4328,
      null,
      2554,
      null,
      2043,
      null,
      2586,
      null,
      4540,
      null,
      4464,
      null,
      3241,
      null,
      2466,
      null,
      1872,
      null,
      534,
      null,
      2068,
      null,
      3578,
      null,
      294,
      null,
      4927,
      null,
      4543,
      null,
      3895,
      null,
      3413,
      null,
      92,
      null,
      1883,
      null,
      4924,
      null,
      2425,
      null,
      3721,
      null,
      964,
      null,
      2399,
      null,
      1853,
      null,
      2892,
      null,
      580,
      null,
      3711,
      null,
      2857,
      null,
      3608,
      null,
      4364,
      null,
      4013,
      null,
      1751,
      null,
      4697,
      null,
      944,
      null,
      1974,
      null,
      4624,
      null,
      638,
      null,
      2187,
      null,
      2891,
      null,
      66,
      null,
      2599,
      null,
      869,
      null,
      4310,
      null,
      2429,
      null,
      2724,
      null,
      582,
      null,
      3068,
      null,
      19,
      null,
      4282,
      null,
      2712,
      null,
      3050,
      null,
      4869,
      null,
      344,
      null,
      2985,
      null,
      2710,
      null,
      1604,
      null,
      2305,
      null,
      3223,
      null,
      117,
      null,
      1088,
      null,
      4652,
      null,
      616,
      null,
      4783,
      null,
      1237,
      null,
      412,
      null,
      2881,
      null,
      654,
      null,
      3285,
      null,
      4890,
      null,
      687,
      null,
      696,
      null,
      4173,
      null,
      4980,
      null,
      3531,
      null,
      923,
      null,
      1556,
      null,
      3496,
      null,
      1761,
      null,
      587,
      null,
      738,
      null,
      2853,
      null,
      3263,
      null,
      3468,
      null,
      708,
      null,
      268,
      null,
      3389,
      null,
      245,
      null,
      2696,
      null,
      3561,
      null,
      3139,
      null,
      1998,
      null,
      3350,
      null,
      831,
      null,
      916,
      null,
      4328,
      null,
      3985,
      null,
      1609,
      null,
      197,
      null,
      29,
      null,
      2039,
      null,
      526,
      null,
      2688,
      null,
      2626,
      null,
      2174,
      null,
      3922,
      null,
      2437,
      null,
      1229,
      null,
      3676,
      null,
      2754,
      null,
      4454,
      null,
      2784,
      null,
      3839,
      null,
      1349,
      null,
      3002,
      null,
      1611,
      null,
      976,
      null,
      816,
      null,
      2024,
      null,
      360,
      null,
      1348,
      null,
      1796,
      null,
      3418,
      null,
      1756,
      null,
      2845,
      null,
      674,
      null,
      2067,
      null,
      2681,
      null,
      4615,
      null,
      588,
      null,
      2936,
      null,
      987,
      null,
      2524,
      null,
      1845,
      null,
      3573,
      null,
      4049,
      null,
      1182,
      null,
      1349,
      null,
      754,
      null,
      3708,
      null,
      910,
      null,
      4239,
      null,
      2354,
      null,
      1651,
      null,
      1618,
      null,
      4768,
      null,
      1309,
      null,
      1260,
      null,
      3020,
      null,
      4928,
      null,
      1943,
      null,
      3428,
      null,
      1996,
      null,
      2582,
      null,
      2534,
      null,
      3498,
      null,
      60,
      null,
      1355,
      null,
      1655,
    ]);
    const result = target(head);
    expect(result).not.toBeNull();
  });
});
