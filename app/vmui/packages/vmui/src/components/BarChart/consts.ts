import {seriesBarsPlugin} from "../../utils/uplot/plugin";
import {barDisp, getBarSeries} from "../../utils/uplot/series";
import {Fill, Stroke} from "../../utils/uplot/types";
import {PaddingSide, Series} from "uplot";


const stroke: Stroke = {
  unit: 3,
  values: (u: { data: number[][]; }) => u.data[2].map((v: number) =>
    v !== 1 ? "#33BB55" : "#F79420"
  ),
};

const fill: Fill = {
  unit: 3,
  values: (u: { data: number[][]; }) => u.data[2].map((v: number) =>
    v !== 1 ? "#33BB55A0" : "#F79420A0"
  ),
};

export const barOptions = {
  height: 500,
  width: 500,
  padding: [null, 0, null, 0] as [top: PaddingSide, right: PaddingSide, bottom: PaddingSide, left: PaddingSide],
  axes: [{ show: false }],
  series: [
    {
      label: "",
      value: (u: uPlot, v: string) => v
    },
    {
      label: "",
      width: 2,
      fill: "#33BB55A0",
      values: (u: uPlot, seriesIdx: number) => {
        const idxs = u.legend.idxs || [];

        if (u.data === null || idxs.length === 0)
          return {"Name": null, "Value": null,};

        const dataIdx = idxs[seriesIdx] || 0;

        const build = u.data[0][dataIdx];
        const duration = u.data[seriesIdx][dataIdx];

        return {"Name": build, "Value": duration};
      }
    },
  ] as Series[],
  plugins: [seriesBarsPlugin(getBarSeries([1], 0, 1, 0, barDisp(stroke, fill)))],
};
