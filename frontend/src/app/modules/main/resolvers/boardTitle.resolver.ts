import { inject } from "@angular/core";
import { ResolveFn } from '@angular/router';
import { BoardModel, selectBoard } from "@entities/board";
import { Store } from '@ngrx/store';
import { filter } from 'rxjs';
import { take } from 'rxjs/operators';


export const boardResolver: ResolveFn<BoardModel> = (route) => {
    const store = inject(Store);
    const boardId = Number(route.paramMap.get('boardId'));
    return store.select(selectBoard.byId(boardId)).pipe(
        filter((board) => board !== null),
        take(1),
    );
};
